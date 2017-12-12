package tournament

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// Tournament holds all data about the tournament
// details about durations etc./competitors/pairings
type Tournament struct {
	Details     Details                  //
	Competitors []competitors.Competitor // a complete slice of all competitors

	Pairings       []pairings.Pairing
	PairingResults []pairings.PairingResult
}

// Details holds Information about the tournament, these are used for serveral calculations and can be seen as a config to the tournament
type Details struct {
	NumberOfParallelGames      int
	MinutesPerGame             int
	MinutesAvailForGroupsPhase int
}

// NewTournamentDetails creates and initializes Details
func NewTournamentDetails(numberOfFields int, minutesPerGame int, minAvail int) *Details {
	d := new(Details)
	d.NumberOfParallelGames = numberOfFields
	d.MinutesPerGame = minutesPerGame
	d.MinutesAvailForGroupsPhase = minAvail
	return d
}

func (t Tournament) getTournamentDetails() Details {
	return t.Details
}

// NewTournament returns a Tournament initialized with the given details
func NewTournament(details Details) Tournament {
	t := Tournament{}
	t.Details = details
	return t
}

func (t *Tournament) setTournamentDetails(td Details) {
	t.Details = td
}

func (t Tournament) build() {

	// calc the tournament plan for the - take in account that no team should play twice in a round ;) and that the last round should be all of one group at once

}
