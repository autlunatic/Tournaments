package tournament

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// T holds all data about the tournament
// details about durations etc./competitors/pairings
const pairingsEntity = "pairings"
const detialsEntity = "details"
const competitorsEntity = "competitors"
const resultsEntity = "results"

type competitorStore struct {
	CompName   string
	DrawNumber int
	GroupID    int
	ID         int
}
type resultStore struct {
	MapIndex int
	result   pairings.Result
}

// T Represents the tournament
type T struct {
	ID          int64
	Details     detail.D        //
	Competitors []competitors.C // a complete slice of all competitors

	Groups         []groups.G
	Pairings       []pairings.P
	PairingResults pairings.Results
	// Plan           [][]pairings.P
	FinalPairings []pairings.P
	PointCalcer   tournamentPoints.TournamentPointCalcer
}

func (t T) getTournamentDetails() detail.D {
	return t.Details
}

func (t *T) setTournamentDetails(td detail.D) {
	t.Details = td
}

// Build calculates the tournament with the given Details and competitors
func (t *T) Build() error {
	// calc the tournament plan for the - take in account that no team should play twice in a round ;) and that the last round should be all of one group at once
	// competitors.CalcRandomDraw(t.Competitors)
	t.Groups, t.Pairings = groups.CalcMostGamesPerCompetitorPlan(t.Competitors, t.Details)
	t.PairingResults = make(map[int]*pairings.Result)
	return nil
}

// GetPairingByID returns the Pairing with the given ID from the tournament
// it returns an Error if the ID is not valid
func (t *T) GetPairingByID(ID int) (pairings.P, error) {
	for i, p := range t.Pairings {
		if p.ID == ID {
			return t.Pairings[i], nil
		}
	}
	for i, p := range t.FinalPairings {
		if p.ID == ID {
			return t.FinalPairings[i], nil
		}
	}

	return pairings.P{}, fmt.Errorf("Invalid Pairing ID")
}

func getCompetitorsKey(c context.Context, id int) *datastore.Key {
	return datastore.NewKey(c, competitorsEntity, strconv.Itoa(id), 0, nil)
}

func getPairingsKey(c context.Context, id int) *datastore.Key {
	return datastore.NewKey(c, pairingsEntity, strconv.Itoa(id), 0, nil)
}

func getFinalPairingsKey(c context.Context, id int) *datastore.Key {
	return datastore.NewKey(c, pairingsEntity+"_F", strconv.Itoa(id), 0, nil)
}

func getResultKey(c context.Context, id int) *datastore.Key {
	return datastore.NewKey(c, resultsEntity, strconv.Itoa(id), 0, nil)
}

func getDetailsKey(c context.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	key := datastore.NewKey(c, detialsEntity, "latest", 0, nil)
	return key
}

// LoadDetails loads the details from datastore to t.Details, if error Details will be se to default
func (t *T) LoadDetails(c context.Context) error {
	var d detail.D
	err := datastore.Get(c, getDetailsKey(c), &d)
	if err != nil || d.FinalistCount == 0 {
		return err
	}
	t.Details = d
	return nil
}

// LoadPairings loads the Pairings from datastore if error the pairings will be set to nil
func (t *T) LoadPairings(c context.Context) {
	q := datastore.NewQuery(pairingsEntity).Order("ID")
	ps := []pairings.P{}
	if _, err := q.GetAll(c, &ps); err != nil {
		log.Errorf(c, "Getting Pairings: %v", err)
		t.Pairings = []pairings.P{}
		log.Errorf(c, "%v", err)
		return
	}
	t.Pairings = ps
	q = datastore.NewQuery(pairingsEntity + "_F").Order("StartTime")
	ps = []pairings.P{}
	if _, err := q.GetAll(c, &ps); err != nil {
		log.Errorf(c, "Getting FinalPairings: %v", err)
		t.FinalPairings = []pairings.P{}
		log.Errorf(c, "%v", err)
		return
	}
	t.FinalPairings = ps
	log.Infof(c, "FinalPairings loaded %v", t.Pairings)
}

// LoadCompetitors loads the name and drawnumber of a saved competitor from the datastore
// saved points are not loaded and must be recalculated by the tournament
func (t *T) LoadCompetitors(c context.Context) {
	t.Competitors = []competitors.C{}
	t.Groups = []groups.G{}
	q := datastore.NewQuery(competitorsEntity).Order("DrawNumber")
	cs := []competitorStore{}
	if _, err := q.GetAll(c, &cs); err != nil {
		log.Errorf(c, "Getting Competitors: %v", err)
		return
	}
	for _, comp := range cs {
		ac := competitors.New(comp.CompName, comp.ID)
		ac.SetDrawNumber(comp.DrawNumber)

		g, errg := groups.GByID(t.Groups, comp.GroupID)
		if errg != nil {
			g = &groups.G{ID: comp.GroupID, Competitors: []competitors.C{}}
			g.AddCompetitor(ac)
			t.Groups = append(t.Groups, *g)
		}
		g.AddCompetitor(ac)

		var err error
		t.Competitors, err = competitors.Add(t.Competitors, ac)
		if err != nil {
			log.Errorf(c, "error loading competitors: %v", err)
			return
		}
	}
	log.Infof(c, "Competitors loaded %v", t.Competitors)
}

// LoadResults loads the results from the datastore
func (t *T) LoadResults(c context.Context) {
	t.PairingResults = make(map[int]*pairings.Result)
	q := datastore.NewQuery(resultsEntity)
	cs := []pairings.Result{}
	var keys []*datastore.Key
	var err error
	keys, err = q.GetAll(c, &cs)
	if err != nil {
		log.Errorf(c, "Getting Results: %v", err)
		return
	}
	for i, comp := range cs {
		id, err := strconv.Atoi(keys[i].StringID())
		if err != nil {
			log.Errorf(c, "Getting Results: %v", err)
			return
		}

		t.PairingResults = pairings.AddResult(t.PairingResults, comp.GamePoints1, comp.GamePoints2, id)
	}
	err = t.RecalcPoints()
	if err != nil {
		log.Errorf(c, "RecalcPoints : %v", err)
		return
	}
	log.Infof(c, "results loaded %v", t.PairingResults)
}

// LoadFromDataStore loads the complete tournament from Datastore
func (t *T) LoadFromDataStore(r *http.Request, w http.ResponseWriter) {
	c := appengine.NewContext(r)
	err := t.LoadDetails(c)
	if err != nil {
		return
	}
	t.LoadPairings(c)
	t.LoadCompetitors(c)
	t.LoadResults(c)
}

// SaveDetails saves tournament Details to the datastore
func (t *T) SaveDetails(c context.Context) error {
	_, err := datastore.Put(c, getDetailsKey(c), &t.Details)
	return err
}

// RecalcPoints calculates the Points for the Groupphase
func (t *T) RecalcPoints() error {
	competitors.ClearPoints(t.Competitors)
	return pairings.AddPointsForResults(t.Competitors, t.Pairings, t.PairingResults, t.PointCalcer)
}

// SavePairings saves the calced pairings to the Database
func (t *T) SavePairings(c context.Context) error {
	for i, pi := range t.Pairings {
		log.Infof(c, "%v", pi)
		_, err := datastore.Put(c, getPairingsKey(c, i), &pi)
		if err != nil {
			return err
		}
	}
	for i, pi := range t.FinalPairings {
		_, err := datastore.Put(c, getFinalPairingsKey(c, i), &pi)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveResults saves the inputted results to the datastore
func (t *T) SaveResults(c context.Context) error {
	for i, pr := range t.PairingResults {
		key := getResultKey(c, i)
		_, err := datastore.Put(c, key, pr)
		if err != nil {
			return err
		}
	}
	log.Infof(c, "results Saved")
	return nil
}

// SaveCompetitors saves the interface to datastore
func (t *T) SaveCompetitors(c context.Context) error {
	for i, ci := range t.Competitors {
		groupID, err := groups.GetGroupIDOfCompetitor(t.Groups, ci.ID())
		if err != nil {
			groupID = -1
		}
		name := competitorStore{ci.Name(), ci.DrawNumber(), groupID, ci.ID()}
		_, err = datastore.Put(c, getCompetitorsKey(c, i), &name)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveToDataStore saves the complete tournament to datastore
func (t *T) SaveToDataStore(r *http.Request, w http.ResponseWriter) {
	c := appengine.NewContext(r)
	if err := t.SaveDetails(c); err != nil {
		log.Errorf(c, "DetailsSave: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := t.SavePairings(c); err != nil {
		log.Errorf(c, "PairingsSave: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := t.SaveCompetitors(c); err != nil {
		log.Errorf(c, "Competitors: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := t.SaveResults(c); err != nil {
		log.Errorf(c, "Results: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// RecalcFinals calculates the Finals from beginning.
// note there must be the finalpairings set from the groupphase this function only recalculates the
// finals from first round to final
func (t *T) RecalcFinals() {
	t.FinalPairings = pairings.RecalcFinals(t.FinalPairings, t.PairingResults, t.PointCalcer, t.Details.NumberOfParallelGames)
	t.SetFinalTimes()
}

// SetFinalTimes calculates the times for the finalPairings
func (t *T) SetFinalTimes() {
	t.FinalPairings = pairings.CalcTimesForFinalPairings(pairings.LatestGameStart(t.Pairings), t.FinalPairings, t.Details)
}

// SetNewCompetitors resets the Competitors if the tournament got a new set of Competitors
func (t *T) SetNewCompetitors(compNames []string) {
	t.Competitors = t.Competitors[:0]
	for _, cn := range compNames {
		var err error
		t.Competitors, err = competitors.AddByName(t.Competitors, cn)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(t.Competitors)
}

// SetNewDetails saves the new details in the tournament.
func (t *T) SetNewDetails(details detail.D) {
	t.Details = details
	t.Details.TournamentStartTime = t.Details.TournamentStartTime.In(time.Now().Location())
	t.Details.FinalsStartTime = t.Details.FinalsStartTime.In(time.Now().Location())
}

// AddMinutesToGroupParingsTime adds the given minutes to all Pairings of the group phase
func (t *T) AddMinutesToGroupParingsTime(minutes int) {
	pairings.AddMinutes(t.Pairings, minutes)
}
