package tournament

type Tournament struct {
	Details     TournamentDetails
	Competitors []Competitor


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

func (t *Tournament) addTournamentDetails(td TournamentDetails) {
	t.Details = td
}
