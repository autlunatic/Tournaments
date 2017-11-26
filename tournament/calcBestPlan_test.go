package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcBestPlan(t *testing.T) {
	competitors := newTestCompetitors(12)
	details := Details{10, 5, 30}

	_, g := calcBestPlan(competitors, details)
	if len(g) != 2 {
		TestingUtils.CheckEquals(2, len(g), "groupCount", t)
	}
}
func TestCalcBestPlanOnly2PerGroupPossible(t *testing.T) {
	competitors := newTestCompetitors(10)
	details := Details{10, 5, 5}

	p, g := calcBestPlan(competitors, details)
	printPlan(p)
	if len(g) != 2 {
		TestingUtils.CheckEquals(5, len(g), "groupCount", t)
	}
}
func TestCalcBestPlanImpossible(t *testing.T) {
	competitors := newTestCompetitors(10)
	details := Details{2, 5, 5}

	p, g := calcBestPlan(competitors, details)
	printPlan(p)
	if len(g) != 2 {
		TestingUtils.CheckEquals(5, len(g), "groupCount", t)
	}
}
