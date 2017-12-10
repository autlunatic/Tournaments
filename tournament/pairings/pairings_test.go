package pairings

import (
	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/tournaments/tournament/competitors"
	"testing"
)

func TestCalcPairings(t *testing.T) {
	benni := competitors.NewCompetitor("Benni")
	dani := competitors.NewCompetitor("Dani")
	c := competitors.Competitors{}
	c.Items = append(c.Items, benni)
	c.Items = append(c.Items, dani)

	pairings := CalcPairings(c.Items, 1)

	TestingUtils.CheckEquals(1, len(pairings), "", t)
	pair := pairings[0]

	if (pair.Competitor1.Name != "Benni") || (pair.Competitor2.Name != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)", pair.Competitor1.Name, pair.Competitor2.Name)
	}
}

func TestCalcPairings3Competitors(t *testing.T) {
	benni := competitors.NewCompetitor("Benni")
	zoe := competitors.NewCompetitor("Zoé")
	dani := competitors.NewCompetitor("Dani")
	c := competitors.Competitors{}
	c.Items = append(c.Items, benni)
	c.Items = append(c.Items, dani)
	c.Items = append(c.Items, zoe)

	pairings := CalcPairings(c.Items, 1)

	TestingUtils.CheckEquals(3, len(pairings), "", t)
	pair := pairings[0]
	if (pair.Competitor1.Name != "Benni") || (pair.Competitor2.Name != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)", pair.Competitor1.Name, pair.Competitor2.Name)
	}
	pair = pairings[1]
	if (pair.Competitor1.Name != "Benni") || (pair.Competitor2.Name != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)", pair.Competitor1.Name, pair.Competitor2.Name)
	}
	pair = pairings[2]
	if (pair.Competitor1.Name != "Dani") || (pair.Competitor2.Name != "Zoé") {
		t.Errorf("first competitor should be Dani (but was %s) and second should be Zoé (but was %s)", pair.Competitor1.Name, pair.Competitor2.Name)
	}
}

func checkPairingDoubles(pairings []Pairing) (msg string) {
	count := 0
	for _, pair := range pairings {
		count = 0
		for _, pair2 := range pairings {
			if isSamePair(pair, pair2) {
				count++
			}
		}
		if count > 1 {
			return pair.Competitor1.Name + " vs. " + pair.Competitor2.Name
		}

	}
	return ""
}
func isSamePair(p Pairing, p2 Pairing) bool {
	return (p.Competitor1.Name == p2.Competitor1.Name) && (p.Competitor2.Name == p2.Competitor2.Name) ||
		(p.Competitor1.Name == p2.Competitor2.Name) && (p.Competitor2.Name == p2.Competitor1.Name)

}

func TestCalcPairings5Competitors(t *testing.T) {
	c := competitors.Competitors{}
	c.Items = append(c.Items,
		competitors.NewCompetitor("Benni"),
		competitors.NewCompetitor("Dani"),
		competitors.NewCompetitor("Zoé"),
		competitors.NewCompetitor("Mona"),
		competitors.NewCompetitor("Andrea"))

	pairings := CalcPairings(c.Items, 1)

	TestingUtils.CheckEquals(10, len(pairings), "", t)
	pair := pairings[0]
	if (pair.Competitor1.Name != "Benni") || (pair.Competitor2.Name != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)", pair.Competitor1.Name, pair.Competitor2.Name)
	}
	pair = pairings[1]
	if (pair.Competitor1.Name != "Mona") || (pair.Competitor2.Name != "Andrea") {
		t.Errorf("first competitor should be Mona (but was %s) and second should be Andrea (but was %s)", pair.Competitor1.Name, pair.Competitor2.Name)
	}
	pair = pairings[2]
	if (pair.Competitor1.Name != "Benni") || (pair.Competitor2.Name != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)", pair.Competitor1.Name, pair.Competitor2.Name)
	}
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same c found! " + msg)
	}
}
func TestCalcPairings12Competitors(t *testing.T) {
	c := competitors.NewTestCompetitors(12)

	pairings := CalcPairings(c.Items, 1)
	TestingUtils.CheckEquals(66, len(pairings), "", t)
	TestingUtils.CheckEquals(6, pairings[33].Round, "Round", t)
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same competitors found! " + msg)
	}
}
