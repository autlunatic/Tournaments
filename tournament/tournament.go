package tournament

import "fmt"

type Tournament struct {
	Details     Details      //
	Competitors []Competitor // a complete slice of all competitors

	Pairings []pairing
}

type Details struct {
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


func NewTournamentDetails(numberOfFields int, minutesPerGame int) *Details {
	d := new(Details)
	d.numberOfParallelGames = numberOfFields
	d.minutesPerGame = minutesPerGame
	return d
}

func (t Tournament) GetTournamentDetails() Details {
	return t.Details
}

func NewTournament(details Details) Tournament {
	t := Tournament{}
	t.Details = details
	return t
}

func (t *Tournament) setTournamentDetails(td Details) {
	t.Details = td
}

func (t Tournament) build() {
	// calc Groups needed given the time you got, the count of Games can be played parallel and the count of competitors -> the goal is that the competitor has as much games as possible

	// calc the tournament plan for the given groupnumber - take in account that no team should play twice in a round ;) and that the last round should be all of one group at once

}
