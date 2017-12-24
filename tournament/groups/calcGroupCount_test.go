package groups

import (
	"testing"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func TestCalcGroupCount(t *testing.T) {
	c := competitors.NewTestCompetitors(12)
	groupCount := CalcGroupCountMaxGames(c, 60, 15, 10)
	TestingUtils.CheckEquals(3, groupCount, "", t)
}
func TestCalcGroupCountImpossible(t *testing.T) {
	c := competitors.NewTestCompetitors(12)
	groupCount := CalcGroupCountMaxGames(c, 10, 15, 10)
	TestingUtils.CheckEquals(0, groupCount, "", t)
}

func TestCalcGroupCountAlInOneGroup(t *testing.T) {

	c := competitors.NewTestCompetitors(12)
	groupCount := CalcGroupCountMaxGames(c, 100000, 15, 10)
	TestingUtils.CheckEquals(1, groupCount, "", t)
}

func TestCalcGroupCount2Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(12)
	groupCount := CalcGroupCountMaxGames(c, 120, 15, 10)
	TestingUtils.CheckEquals(2, groupCount, "", t)
}
func TestCalcGroupCount3Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(30)
	groupCount := CalcGroupCountMaxGames(c, 136, 15, 10)
	TestingUtils.CheckEquals(5, groupCount, "", t)
}

func TestCalcGroupCount9Competitors(t *testing.T) {
	c := competitors.NewTestCompetitors(9)
	groupCount := CalcGroupCountMaxGames(c, 80, 10, 2)
	TestingUtils.CheckEquals(2, groupCount, "", t)
	groupCount = CalcGroupCountMaxGames(c, 60, 10, 3)
	TestingUtils.CheckEquals(2, groupCount, "3 prallel 60 mins", t)
	groupCount = CalcGroupCountMaxGames(c, 50, 10, 4)
	TestingUtils.CheckEquals(2, groupCount, "4 prallel 50 mins", t)
	groupCount = CalcGroupCountMaxGames(c, 50, 10, 10)
	TestingUtils.CheckEquals(2, groupCount, "10 parallel 50 mins", t)
	groupCount = CalcGroupCountMaxGames(c, 40, 10, 10)
	TestingUtils.CheckEquals(3, groupCount, "10 parallel 40 mins", t)
}
