package tournament

type Tournament struct {
	details      TournamentDetails
	contributors []Contributor
}

func (t Tournament) GetTournamentDetails() TournamentDetails {
	return t.details
}

func NewTournament(details TournamentDetails) Tournament {
	t := Tournament{details, []Contributor{}}
	return t
}

func (t *Tournament) addTournamentDetails(td TournamentDetails) {
	t.details = td
}
