package groups

import (
	"testing"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournament"
)

func TestCalcBestPlan(t *testing.T) {
	competitors.Items = competitors.NewTestCompetitors(12)
	details := tournament.Details{NumberOfParallelGames: 10, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 30}

	_, g := calcMostGamesPerCompetitorPlan(competitors.Items, details)
	TestingUtils.CheckEquals(2, len(g), "groupCount", t)
}
func TestCalcBestPlanOnly2PerGroupPossible(t *testing.T) {
	competitors.Items = competitors.NewTestCompetitors(10)
	details := tournament.Details{NumberOfParallelGames: 10, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 5}

	_, g := calcMostGamesPerCompetitorPlan(competitors.Items, details)
	if len(g) != 2 {
		TestingUtils.CheckEquals(5, len(g), "groupCount", t)
	}
}
func TestCalcBestPlanImpossible(t *testing.T) {
	competitors.Items = competitors.NewTestCompetitors(10)
	details := tournament.Details{NumberOfParallelGames: 2, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 5}

	p, _ := calcMostGamesPerCompetitorPlan(competitors.Items, details)
	if len(p) != 0 {
		t.Error("plan should be empty because it is not possible to do a tournament with given values")
	}
}
