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

func (t *Tournament) AddContributor(contributor Contributor) {
	t.contributors = append(t.contributors, contributor)
}

func newTestTournament() Tournament {
	td := NewTournamentDetails(10, 2, 1)
	tournament := NewTournament(*td)
	return tournament
}

func TestAddContributor(t *testing.T) {
	var tournament Tournament
	tournament = newTestTournament()
	tournament.AddContributor(Contributor{"Benni"})
	if len(tournament.contributors) != 1 {
		t.Error("Should be one contributor")
	}
	if tournament.contributors[0].name != "Benni" {
		t.Error("first contributors name should be Benni")
	}

}
