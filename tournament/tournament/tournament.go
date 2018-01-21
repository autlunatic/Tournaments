package tournament

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// Tournament holds all data about the tournament
// details about durations etc./competitors/pairings
type Tournament struct {
	Details     detail.D        //
	Competitors []competitors.C // a complete slice of all competitors

	Groups         []groups.G
	Pairings       []pairings.P
	PairingResults []pairings.Results
}

func (t Tournament) getTournamentDetails() detail.D {
	return t.Details
}

// NewTournament returns a Tournament initialized with the given details
func NewTournament(details detail.D) Tournament {
	t := Tournament{}
	t.Details = details
	return t
}

func (t *Tournament) setTournamentDetails(td detail.D) {
	t.Details = td
}

func (t Tournament) build() {

	// calc the tournament plan for the - take in account that no team should play twice in a round ;) and that the last round should be all of one group at once

}
