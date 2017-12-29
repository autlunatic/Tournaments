package pairings

import (
	"testing"

	"fmt"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func TestCalcPairings(t *testing.T) {

	benni := competitors.New("Benni", 0)
	dani := competitors.New("Dani", 1)
	cs := make([]competitors.Competitor, 2)
	cs[0] = benni
	cs[1] = dani

	pairings, _ := CalcPairings(cs, 1)

	TestingUtils.CheckEquals(1, len(pairings), "", t)
	pair := pairings[0]

	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
}
func TestCalcPairingsEmptyCompetitorsShouldNotPanic(t *testing.T) {
	var cs []competitors.Competitor
	_, err := CalcPairings(cs, 1)
	if err == nil {
		t.Error("expected Error didnt return!")
	}
}

func TestCalcPairings3Competitors(t *testing.T) {
	benni := competitors.New("Benni", 0)
	dani := competitors.New("Dani", 1)
	zoe := competitors.New("Zoé", 2)
	cs := make([]competitors.Competitor, 3)
	cs[0] = benni
	cs[1] = dani
	cs[2] = zoe

	pairings, _ := CalcPairings(cs, 1)

	TestingUtils.CheckEquals(3, len(pairings), "", t)
	pair := pairings[0]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	pair = pairings[1]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())

	}
	pair = pairings[2]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Dani") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Zoé") {
		t.Errorf("first competitor should be Dani (but was %s) and second should be Zoé (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
}

func checkPairingDoubles(pairings []Pairing) (msg string) {
	var count int
	for _, pair := range pairings {
		count = 0
		for _, pair2 := range pairings {
			if isSamePair(pair, pair2) {
				count++
			}
		}
		if count > 1 {
			return fmt.Sprintf("CompetitorIds: %v vs. %v", pair.Competitor1ID, pair.Competitor2ID)
		}

	}
	return ""
}
func isSamePair(p Pairing, p2 Pairing) bool {
	return (p.Competitor1ID == p2.Competitor1ID) && (p.Competitor2ID == p2.Competitor2ID) ||
		(p.Competitor1ID == p2.Competitor2ID) && (p.Competitor2ID == p2.Competitor1ID)

}

func TestCalcPairings5Competitors(t *testing.T) {
	var cs []competitors.Competitor
	cs = append(cs,
		competitors.New("Benni", 0),
		competitors.New("Dani", 1),
		competitors.New("Zoé", 2),
		competitors.New("Mona", 3),
		competitors.New("Andrea", 4))

	pairings, _ := CalcPairings(cs, 1)

	TestingUtils.CheckEquals(10, len(pairings), "", t)
	pair := pairings[0]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	pair = pairings[1]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Mona") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Andrea") {
		t.Errorf("first competitor should be Mona (but was %s) and second should be Andrea (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	pair = pairings[2]
	if competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni" || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same c found! " + msg)
	}
}
func TestCalcPairings12Competitors(t *testing.T) {
	cs := competitors.NewTestCompetitors(12)

	pairings, _ := CalcPairings(cs, 1)
	TestingUtils.CheckEquals(66, len(pairings), "", t)
	TestingUtils.CheckEquals(6, pairings[33].Round, "Round", t)
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same competitors found! " + msg)
	}
}
