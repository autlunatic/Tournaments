package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcGroupCount(t *testing.T) {
	competitors := newTestCompetitors(12)
	groupCount := calcGroupCountMaxGames(competitors, 60, 15, 10)
	TestingUtils.CheckEquals(3, groupCount, "", t)
}
func TestCalcGroupCountImpossible(t *testing.T) {
	competitors := newTestCompetitors(12)
	groupCount := calcGroupCountMaxGames(competitors, 10, 15, 10)
	TestingUtils.CheckEquals(0, groupCount, "", t)
}

func TestCalcGroupCountAlInOneGroup(t *testing.T) {
	competitors := newTestCompetitors(12)
	groupCount := calcGroupCountMaxGames(competitors, 100000, 15, 10)
	TestingUtils.CheckEquals(1, groupCount, "", t)
}

func TestCalcGroupCount2Groups(t *testing.T) {
	competitors := newTestCompetitors(12)
	groupCount := calcGroupCountMaxGames(competitors, 120, 15, 10)
	TestingUtils.CheckEquals(2, groupCount, "", t)
}
func TestCalcGroupCount3Groups(t *testing.T) {
	competitors := newTestCompetitors(30)
	groupCount := calcGroupCountMaxGames(competitors, 136, 15, 10)
	TestingUtils.CheckEquals(3, groupCount, "", t)
}

func TestCalcGroupCount2Competitors(t *testing.T) {
	competitors := newTestCompetitors(2)
	groupCount := calcGroupCountMaxGames(competitors, 150, 15, 10)
	TestingUtils.CheckEquals(1, groupCount, "", t)
}
