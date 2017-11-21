package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcDurationInterface(t *testing.T) {
	td := NewTournamentDetails(1, 2, 3)
	var result Tournament
	result = NewTournament(*td)

	TestingUtils.CheckEquals(3, InMinutes(result), "",t)
}

func (t *Tournament) AddCompetitor(c Competitor) {
	t.competitors = append(t.competitors, c)
}

func newTestTournament() Tournament {
	td := NewTournamentDetails(10, 2, 1)
	tournament := NewTournament(*td)
	return tournament
}

func TestAddCompetitor(t *testing.T) {
	var tournament Tournament
	tournament = newTestTournament()
	tournament.AddCompetitor(Competitor{"Benni"})
	if len(tournament.competitors) != 1 {
		t.Error("Should be one competitor")
	}
	if tournament.competitors[0].name != "Benni" {
		t.Error("first competitors name should be Benni")
	}

}
