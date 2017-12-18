package tournament

import (
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func (t *Tournament) AddCompetitor(c competitors.Competitor) {
	t.Competitors = append(t.Competitors, c)
}

func newTestTournament() Tournament {
	td := NewTournamentDetails(2, 1, 0)
	tournament := NewTournament(*td)
	return tournament
}

func TestAddCompetitor(t *testing.T) {
	tournament := newTestTournament()
	tournament.AddCompetitor(competitors.NewCompetitor("Benni", 0))
	if len(tournament.Competitors) != 1 {
		t.Error("Should be one competitor")
	}
	if tournament.Competitors[0].Name() != "Benni" {
		t.Error("first competitors name should be Benni")
	}

}
