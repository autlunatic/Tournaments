package tournament

type Tournament struct {
	details     TournamentDetails
	competitors []Competitor
}

func (t Tournament) GetTournamentDetails() TournamentDetails {
	return t.details
}

func NewTournament(details TournamentDetails) Tournament {
	t := Tournament{details, []Competitor{}}
	return t
}

func (t *Tournament) addTournamentDetails(td TournamentDetails) {
	t.details = td
}
