package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"strconv"
	"testing"
)

func TestCalcPairings(t *testing.T) {
	benni := newCompetitor("Benni")
	dani := newCompetitor("Dani")
	competitors := Competitors{}
	competitors.items = append(competitors.items, benni)
	competitors.items = append(competitors.items, dani)

	pairings := calcPairings(competitors)

	TestingUtils.CheckEquals(1, len(pairings), "", t)
	pair := pairings[0]

	if (pair.competitor1.name != "Benni") || (pair.competitor2.name != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
}

func TestCalcPairings3Competitors(t *testing.T) {
	benni := newCompetitor("Benni")
	zoe := newCompetitor("Zoé")
	dani := newCompetitor("Dani")
	competitors := Competitors{}
	competitors.items = append(competitors.items, benni)
	competitors.items = append(competitors.items, dani)
	competitors.items = append(competitors.items, zoe)

	pairings := calcPairings(competitors)

	TestingUtils.CheckEquals(3, len(pairings), "", t)
	pair := pairings[0]
	if (pair.competitor1.name != "Benni") || (pair.competitor2.name != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
	pair = pairings[1]
	if (pair.competitor1.name != "Benni") || (pair.competitor2.name != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
	pair = pairings[2]
	if (pair.competitor1.name != "Dani") || (pair.competitor2.name != "Zoé") {
		t.Errorf("first competitor should be Dani (but was %s) and second should be Zoé (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
}

func checkPairingDoubles(pairings []pairing) (msg string) {
	count := 0
	for _, pair := range pairings {
		count = 0
		for _, pair2 := range pairings {
			if isSamePair(pair, pair2) {
				count++
			}
		}
		if count > 1 {
			return pair.competitor1.name + " vs. " + pair.competitor2.name
		}

	}
	return ""
}
func isSamePair(p pairing, p2 pairing) bool {
	return (p.competitor1.name == p2.competitor1.name) && (p.competitor2.name == p2.competitor2.name) ||
		(p.competitor1.name == p2.competitor2.name) && (p.competitor2.name == p2.competitor1.name)

}

func TestCalcPairings5Competitors(t *testing.T) {
	competitors := Competitors{}
	competitors.items = append(competitors.items,
		newCompetitor("Benni"),
		newCompetitor("Dani"),
		newCompetitor("Zoé"),
		newCompetitor("Mona"),
		newCompetitor("Andrea"))

	pairings := calcPairings(competitors)

	TestingUtils.CheckEquals(10, len(pairings), "", t)
	pair := pairings[0]
	if (pair.competitor1.name != "Benni") || (pair.competitor2.name != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
	pair = pairings[1]
	if (pair.competitor1.name != "Mona") || (pair.competitor2.name != "Andrea") {
		t.Errorf("first competitor should be Mona (but was %s) and second should be Andrea (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
	pair = pairings[2]
	if (pair.competitor1.name != "Benni") || (pair.competitor2.name != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same Competitors found! " + msg)
	}
}
func TestCalcPairings12Competitors(t *testing.T) {
	competitors := Competitors{}
	for i := 1; i < 13; i++ {
		competitors.items = append(competitors.items, newCompetitor(strconv.Itoa(i)))
	}

	pairings := calcPairings(competitors)
	TestingUtils.CheckEquals(66, len(pairings), "", t)
	TestingUtils.CheckEquals(6, pairings[33].round, "Round", t)
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same Competitors found! " + msg)
	}
}
