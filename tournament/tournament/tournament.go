package tournament

import (
	"fmt"

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
	return pairings.P{}, fmt.Errorf("Invalid Pairing ID")
}
