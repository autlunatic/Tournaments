package tournament

import (
	"testing"
)

func (t *Tournament) AddCompetitor(c Competitor) {
	t.Competitors = append(t.Competitors, c)
}

func newTestTournament() Tournament {
	td := newTournamentDetails(2, 1)
	tournament := NewTournament(*td)
	return tournament
}

func TestAddCompetitor(t *testing.T) {
	tournament := newTestTournament()
	tournament.AddCompetitor(newCompetitor("Benni"))
	if len(tournament.Competitors) != 1 {
		t.Error("Should be one competitor")
	}
	if tournament.Competitors[0].name != "Benni" {
		t.Error("first competitors name should be Benni")
	}

}
