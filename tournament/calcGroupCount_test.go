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
	TestingUtils.CheckEquals(5, groupCount, "", t)
}

func TestCalcGroupCount9Competitors(t *testing.T) {
	competitors := newTestCompetitors(9)
	groupCount := calcGroupCountMaxGames(competitors, 80, 10, 2)
	TestingUtils.CheckEquals(2, groupCount, "", t)
	groupCount = calcGroupCountMaxGames(competitors, 60, 10, 3)
	TestingUtils.CheckEquals(2, groupCount, "3 prallel 60 mins", t)
	groupCount = calcGroupCountMaxGames(competitors, 50, 10, 4)
	TestingUtils.CheckEquals(2, groupCount, "4 prallel 50 mins", t)
	groupCount = calcGroupCountMaxGames(competitors, 50, 10, 10)
	TestingUtils.CheckEquals(2, groupCount, "10 parallel 50 mins", t)
	groupCount = calcGroupCountMaxGames(competitors, 40, 10, 10)
	TestingUtils.CheckEquals(3, groupCount, "10 parallel 40 mins", t)
}
