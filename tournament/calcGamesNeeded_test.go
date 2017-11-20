package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcGroup(t *testing.T) {
	result := calcGroupGamesNeeded(1)
	TestingUtils.CheckEquals(0, result,"", t)
}

func TestCalcGroupTwoTeams(t *testing.T) {
	result := calcGroupGamesNeeded(2)
	TestingUtils.CheckEquals(2, result, "",t)
}

func TestCalcGroup3Teams(t *testing.T) {
	result := calcGroupGamesNeeded(3)
	TestingUtils.CheckEquals(3, result, "",t)
}

func TestCalcGroup4Teams(t *testing.T) {
	result := calcGroupGamesNeeded(4)
	TestingUtils.CheckEquals(6, result, "",t)
}
