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
	gs := []Group{{id: 1, competitors: c}}
	plan := calcPlan(gs, 1)
	if competitors.GetCompetitor(plan[0][0].Competitor1ID).Name() != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + competitors.GetCompetitor(plan[0][0].Competitor1ID).Name())
	}
	TestingUtils.CheckEquals(6, len(plan), "roundcount should be 6", t)
}

func TestCalcTournamentPlan2Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(8)
	groups := []Group{
		{id: 1, competitors: c[0:4]},
		{id: 2, competitors: c[4:8]}}
	plan := calcPlan(groups, 2)
	if competitors.GetCompetitor(plan[0][0].Competitor1ID).Name() != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + competitors.GetCompetitor(plan[0][0].Competitor1ID).Name())
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
		{id: 1, competitors: c[0:3]},
		{id: 2, competitors: c[3:6]},
		{id: 3, competitors: c[6:9]},
		{id: 4, competitors: c[9:11]},
	}
	plan := calcPlan(groups, 2)
	if competitors.GetCompetitor(plan[0][0].Competitor1ID).Name() != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + competitors.GetCompetitor(plan[0][0].Competitor1ID).Name())
	}
	TestingUtils.CheckEquals(5, len(plan), "roundcount", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentPlan2Groups6_oneCouldNotPlayToTimesInOneRow(t *testing.T) {
	c := competitors.NewTestCompetitors(6)
	groups := []Group{
		{id: 1, competitors: c[0:3]},
		{id: 2, competitors: c[3:6]},
	}
	plan := calcPlan(groups, 10)
	if competitors.GetCompetitor(plan[0][0].Competitor1ID).Name() != "1" {
		t.Error("competitor 0 0 should be named 1 but was " + competitors.GetCompetitor(plan[0][0].Competitor1ID).Name())
	}
	TestingUtils.CheckEquals(3, len(plan), "roundcount should be 6", t)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
}

func TestCalcTournamentLastRoundOfGroupShouldNotBeSplittedOverFieldRounds(t *testing.T) {
	c := competitors.NewTestCompetitors(7)
	groups := []Group{
		{id: 1, competitors: c[0:3]},
		{id: 2, competitors: c[3:7]},
	}
	plan := calcPlan(groups, 2)
	if ok, msg := checkNoCompetitorPlaysTwiceInARound(plan); !ok {
		t.Error(msg)
	}
	if len(plan[3]) != 1 {
		t.Errorf("in Fieldround 3 should only be 1 game but was %d", len(plan[3]))
	}
	mp := plan[4][0]
	if competitors.GetCompetitor(mp.Competitor1ID).Name() != "4" || competitors.GetCompetitor(mp.Competitor2ID).Name() != "7" {
		t.Error("4v7 should be played in fieldround 4")
	}
}

func checkNoCompetitorPlaysTwiceInARound(p [][]pairings.Pairing) (bool, string) {
	for _, round := range p {
		for i, p1 := range round {
			for j, p2 := range round {
				if (i != j) &&
					(p1.Competitor1ID == p2.Competitor1ID ||
						p1.Competitor2ID == p2.Competitor2ID ||
						p2.Competitor2ID == p1.Competitor1ID ||
						p2.Competitor1ID == p1.Competitor2ID) {

					return false, fmt.Sprintf("%v;%v", p1, p2)
				}
			}
		}
	}
	return true, ""
}
