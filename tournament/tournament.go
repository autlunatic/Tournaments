package tournament

type Tournament struct {
	Details     TournamentDetails //
	Competitors []Competitor      // a complete slice of all competitors

	Pairings []pairing
}


func (t Tournament) GetTournamentDetails() TournamentDetails {
	return t.Details
}

func NewTournament(details TournamentDetails) Tournament {
	t := Tournament{}
	t.Details = details
	return t
}

func (t *Tournament) setTournamentDetails(td TournamentDetails) {
	t.Details = td
}

func (t Tournament) build() {
	// calc Groups needed given the time you got, the count of Games can be played parallel and the count of competitors -> the goal is that the competitor has as much games as possible

	// calc the tournament plan for the given groupnumber - take in account that no team should play twice in a round ;) and that the last round should be all of one group at once

}
