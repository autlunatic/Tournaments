package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcDurationInterface(t *testing.T) {
	td := NewTournamentDetails(1, 2, 3)
	result := NewTournament(*td)

	TestingUtils.CheckEquals(3, InMinutes(result), "", t)
}

func (t *Tournament) AddCompetitor(c Competitor) {
	t.Competitors = append(t.Competitors, c)
}

func newTestTournament() Tournament {
	td := NewTournamentDetails(10, 2, 1)
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
