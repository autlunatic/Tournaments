package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcTournamentPlan(t *testing.T) {
	c := newTestCompetitors(4)
	gs := []group{{1, Competitors{c.items[0:4]}}}
	plan := calcPlan(gs, 1)
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
	plan := calcPlan(groups, 2)
	if plan[0][0].competitor1.name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].competitor1.name)
	}
	TestingUtils.CheckEquals(6, len(plan), "roundcount", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentPlan4Groups11(t *testing.T) {
	c := newTestCompetitors(11)
	groups := []group{
		{1, Competitors{c.items[0:3]}},
		{2, Competitors{c.items[3:6]}},
		{3, Competitors{c.items[6:9]}},
		{4, Competitors{c.items[9:11]}},
	}
	plan := calcPlan(groups, 2)
	if plan[0][0].competitor1.name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].competitor1.name)
	}
	TestingUtils.CheckEquals(5, len(plan), "roundcount", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentPlan2Groups6_oneCouldNotPlayToTimesInOneRow(t *testing.T) {
	c := newTestCompetitors(6)
	groups := []group{
		{1, Competitors{c.items[0:3]}},
		{2, Competitors{c.items[3:6]}},
	}
	plan := calcPlan(groups, 10)
	if plan[0][0].competitor1.name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].competitor1.name)
	}
	TestingUtils.CheckEquals(3, len(plan), "roundcount should be 6", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentLastRoundOfGroupShouldNotBeSplittedOverFieldRounds(t *testing.T) {
	c := newTestCompetitors(7)
	groups := []group{
		{1, Competitors{c.items[0:3]}},

		{2, Competitors{c.items[3:7]}},
	}
	plan := calcPlan(groups, 2)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
	if len(plan[3]) != 1 {
		t.Errorf("in Fieldround 3 should only be 1 game but was %d", len(plan[3]))
	}
	mp := plan[4][0]
	if mp.competitor1.name != "4" || mp.competitor2.name != "7" {
		t.Error("4v7 should be played in fieldround 4")
	}
}

func checkNoCompetitorPlaysTwiceInARound(p [][]pairing) (bool, string) {
	for _, round := range p {
		for i, p1 := range round {
			for j, p2 := range round {
				if (i != j) &&
					(p1.competitor1 == p2.competitor1 ||
						p1.competitor2 == p2.competitor2 ||
						p2.competitor2 == p1.competitor1||
						p2.competitor1 == p1.competitor2) {
					return false, p1.toString() + "; " + p2.toString()
				}
			}
		}
	}
	return true, ""
}
