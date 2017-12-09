package tournament

import "fmt"

// Tournament holds all data about the tournament
// details about durations etc./competitors/pairings
type Tournament struct {
	Details     details      //
	Competitors []Competitor // a complete slice of all competitors

	Pairings []pairing
}

type details struct {
	numberOfParallelGames      int
	minutesPerGame             int
	minutesAvailForGroupsPhase int
}

func printPlan(plan [][]pairing) {
	for fieldRound, r := range plan {
		for field, p := range r {
			fmt.Println(fmt.Sprintf("fieldround: %d; field: %d; pairing:", fieldRound, field) + p.toString())
		}
	}
}


func newTournamentDetails(numberOfFields int, minutesPerGame int) *details {
	d := new(details)
	d.numberOfParallelGames = numberOfFields
	d.minutesPerGame = minutesPerGame
	return d
}

func (t Tournament) getTournamentDetails() details {
	return t.Details
}

// NewTournament returns a Tournament initialized with the given details
func NewTournament(details details) Tournament {
	t := Tournament{}
	t.Details = details
	return t
}

func (t *Tournament) setTournamentDetails(td details) {
	t.Details = td
}

func (t Tournament) build() {
	// calc Groups needed given the time you got, the count of Games can be played parallel and the count of competitors -> the goal is that the competitor has as much games as possible

	// calc the tournament plan for the given groupnumber - take in account that no team should play twice in a round ;) and that the last round should be all of one group at once

}
