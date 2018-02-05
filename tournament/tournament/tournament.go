package tournament

import (
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
}

func (t T) getTournamentDetails() detail.D {
	return t.Details
}

// NewTournament returns a Tournament initialized with the given details
func NewTournament(details detail.D) T {
	t := T{}
	t.Details = details
	return t
}

func (t *T) setTournamentDetails(td detail.D) {
	t.Details = td
}

// Build calculates the tournament with the given Details and competitors
func (t *T) Build() error {
	// calc the tournament plan for the - take in account that no team should play twice in a round ;) and that the last round should be all of one group at once
	t.Details = detail.D{
		MinutesAvailForGroupsPhase: 90,
		MinutesPerGame:             15,
		NumberOfParallelGames:      4,
	}
	t.Competitors = competitors.NewTestCompetitors(9)
	t.Plan, t.Groups, t.Pairings = groups.CalcMostGamesPerCompetitorPlan(t.Competitors, t.Details)
	return nil
}
