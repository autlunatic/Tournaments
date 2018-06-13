package tournament

import (
	"fmt"
	"log"

	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// T holds all data about the tournament
// details about durations etc./competitors/pairings
type T struct {
	ID          int64
	Details     detail.D        //
	Competitors []competitors.C // a complete slice of all competitors

	Groups         []groups.G
	Pairings       []pairings.P
	PairingResults pairings.Results
	Plan           [][]pairings.P
	FinalPairings  []pairings.P
	PointCalcer    tournamentPoints.TournamentPointCalcer
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
	t.Plan, t.Groups, t.Pairings = groups.CalcMostGamesPerCompetitorPlan(t.Competitors, t.Details)
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

// RecalcFinals calculates the Finals from beginning.
// note there must be the finalpairings set from the groupphase this function only recalculates the
// finals from first round to final
func (t *T) RecalcFinals() {
	t.FinalPairings = pairings.RecalcFinals(t.FinalPairings, t.PairingResults, t.PointCalcer)
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
		fmt.Println(cn, "added")
		var err error
		t.Competitors, err = competitors.AddByName(t.Competitors, cn)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Println(t.Competitors)
}
