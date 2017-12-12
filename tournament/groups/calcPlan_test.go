package groups

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

func TestCalcTournamentPlan(t *testing.T) {
	c := competitors.NewTestCompetitors(4)

	gs := []Group{{1, competitors: competitors.Competitors{c.Items[0:4]}}}
	plan := calcPlan(gs, 1)
	if plan[0][0].Competitor1.Name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].Competitor1.Name)
	}
	TestingUtils.CheckEquals(6, len(plan), "roundcount should be 6", t)
}

func TestCalcTournamentPlan2Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(8)
	groups := []Group{
		{1, competitors.Competitors{c.Items[0:4]}},
		{2, competitors.Competitors{c.Items[4:8]}}}
	plan := calcPlan(groups, 2)
	if plan[0][0].Competitor1.Name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].Competitor1.Name)
	}
	if plan[1][0].ID != 2 {
		t.Error("id of the first game in second round should be 2" + strconv.Itoa(plan[1][0].ID))
	}
	TestingUtils.CheckEquals(6, len(plan), "roundcount", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentPlan4Groups11(t *testing.T) {
	c := competitors.NewTestCompetitors(11)
	groups := []Group{
		{1, competitors.Competitors{c.Items[0:3]}},
		{2, competitors.Competitors{c.Items[3:6]}},
		{3, competitors.Competitors{c.Items[6:9]}},
		{4, competitors.Competitors{c.Items[9:11]}},
	}
	plan := calcPlan(groups, 2)
	if plan[0][0].Competitor1.Name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].Competitor1.Name)
	}
	TestingUtils.CheckEquals(5, len(plan), "roundcount", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentPlan2Groups6_oneCouldNotPlayToTimesInOneRow(t *testing.T) {
	c := competitors.NewTestCompetitors(6)
	groups := []Group{
		{1, competitors.Competitors{c.Items[0:3]}},
		{2, competitors.Competitors{c.Items[3:6]}},
	}
	plan := calcPlan(groups, 10)
	if plan[0][0].Competitor1.Name != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + plan[0][0].Competitor1.Name)
	}
	TestingUtils.CheckEquals(3, len(plan), "roundcount should be 6", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentLastRoundOfGroupShouldNotBeSplittedOverFieldRounds(t *testing.T) {
	c := competitors.NewTestCompetitors(7)
	groups := []Group{
		{1, competitors.Competitors{c.Items[0:3]}},

		{2, competitors.Competitors{c.Items[3:7]}},
	}
	plan := calcPlan(groups, 2)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
	if len(plan[3]) != 1 {
		t.Errorf("in Fieldround 3 should only be 1 game but was %d", len(plan[3]))
	}
	mp := plan[4][0]
	if mp.Competitor1.Name != "4" || mp.Competitor2.Name != "7" {
		t.Error("4v7 should be played in fieldround 4")
	}
}

func checkNoCompetitorPlaysTwiceInARound(p [][]pairings.Pairing) (bool, string) {
	for _, round := range p {
		for i, p1 := range round {
			for j, p2 := range round {
				if (i != j) &&
					(p1.Competitor1 == p2.Competitor1 ||
						p1.Competitor2 == p2.Competitor2 ||
						p2.Competitor2 == p1.Competitor1 ||
						p2.Competitor1 == p1.Competitor2) {
					return false, fmt.Sprintf("%v;%v", p1, p2)
				}
			}
		}
	}
	return true, ""
}
