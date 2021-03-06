package groups

import (
	"testing"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
)

func TestCalcBestPlan(t *testing.T) {
	c := competitors.NewTestCompetitors(12)
	details := detail.D{NumberOfParallelGames: 10, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 30}

	g, _ := CalcMostGamesPerCompetitorPlan(c, details)
	TestingUtils.CheckEquals(2, len(g), "groupCount", t)
}
func TestCalcBestPlanOnly2PerGroupPossible(t *testing.T) {
	c := competitors.NewTestCompetitors(10)
	details := detail.D{NumberOfParallelGames: 10, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 5}

	g, _ := CalcMostGamesPerCompetitorPlan(c, details)
	if len(g) != 2 {
		TestingUtils.CheckEquals(5, len(g), "groupCount", t)
	}
}
func TestCalcBestPlanImpossible(t *testing.T) {
	c := competitors.NewTestCompetitors(10)
	details := detail.D{NumberOfParallelGames: 2, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 5}

	_, p := CalcMostGamesPerCompetitorPlan(c, details)
	if len(p) != 0 {
		t.Error("plan should be empty because it is not possible to do a tournament with given values")
	}
}
