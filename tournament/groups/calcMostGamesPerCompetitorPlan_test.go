package groups

import (
	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/tournaments/tournament/competitors"
	"github.com/autlunatic/tournaments/tournament/tournament"
	"testing"
)

func TestCalcBestPlan(t *testing.T) {
	c := competitors.NewTestCompetitors(12)
	details := tournament.Details{10, 5, 30}

	_, g := calcMostGamesPerCompetitorPlan(c, details)
	TestingUtils.CheckEquals(2, len(g), "groupCount", t)
}
func TestCalcBestPlanOnly2PerGroupPossible(t *testing.T) {
	c := competitors.NewTestCompetitors(10)
	details := tournament.Details{10, 5, 5}

	_, g := calcMostGamesPerCompetitorPlan(c, details)
	if len(g) != 2 {
		TestingUtils.CheckEquals(5, len(g), "groupCount", t)
	}
}
func TestCalcBestPlanImpossible(t *testing.T) {
	c := competitors.NewTestCompetitors(10)
	details := tournament.Details{2, 5, 5}

	p, _ := calcMostGamesPerCompetitorPlan(c, details)
	if len(p) != 0 {
		t.Error("plan should be empty because it is not possible to do a tournament with given values")
	}
}
