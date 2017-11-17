package tournament

import (
	"testing"
	"github.com/autlunatic/TestingUtils"
)

func TestCalcDurationInterface(t *testing.T) {
	td := NewTournamentDetails(1,2,3)
    var result Tournament
	result = NewTournament(*td)

	lol := InMinutes(result)
	TestingUtils.CheckEquals(3, lol, t)
}

func (t *Tournament) AddContributor(contributor Contributor){
	t.contributors = append(t.contributors, contributor)
}


func newTestTournament() Tournament{
	td := NewTournamentDetails(10,2,1)
	tournament := NewTournament(*td)
	return tournament
}

func TestAddContributor(t *testing.T){
	var tournament Tournament
	tournament = newTestTournament()
	tournament.AddContributor(Contributor{"Benni"})
	if len(tournament.contributors) != 1{
		t.Error()
	}

}