package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcTournamentPlan(t *testing.T) {
	c := newTestCompetitors(4)
	gs := []group{{1, Competitors{c.items[0:4]}}}
	plan := calcTournamentPlan(gs, 1)
	if plan[0][0].competitor1.name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].competitor1.name)
	}
	TestingUtils.CheckEquals(6, len(plan), "roundcount should be 6", t)
}

func TestCalcTournamentPlan2Groups(t *testing.T) {
	c := newTestCompetitors(8)
	groups := []group{
		{1, Competitors{c.items[0:4]}},
		{2, Competitors{c.items[4:8]}}}
	plan := calcTournamentPlan(groups, 2)
	if plan[0][0].competitor1.name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].competitor1.name)
	}
	TestingUtils.CheckEquals(6, len(plan), "roundcount", t)
}

func TestCalcTournamentPlan4Groups11(t *testing.T) {
	c := newTestCompetitors(11)
	groups := []group{
		{1, Competitors{c.items[0:3]}},
		{2, Competitors{c.items[3:6]}},
		{3, Competitors{c.items[6:9]}},
		{4, Competitors{c.items[9:11]}},
	}
	plan := calcTournamentPlan(groups, 2)
	if plan[0][0].competitor1.name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].competitor1.name)
	}
	TestingUtils.CheckEquals(5, len(plan), "roundcount", t)
}

func TestCalcTournamentPlan2Groups6_oneCouldNotPlayToTimesInOneRow(t *testing.T) {
	c := newTestCompetitors(6)
	groups := []group{
		{1, Competitors{c.items[0:3]}},
		{2, Competitors{c.items[3:6]}},
	}
	plan := calcTournamentPlan(groups, 10)
	if plan[0][0].competitor1.name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].competitor1.name)
	}
	TestingUtils.CheckEquals(3, len(plan), "roundcount should be 6", t)
}
