package tournament

import (
	"fmt"
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcPairings(t *testing.T) {
	benni := Competitor{"Benni"}
	dani := Competitor{"Dani"}
	competitors := Competitors{}
	competitors.items = append(competitors.items, benni)
	competitors.items = append(competitors.items, dani)

	pairings := calcPairingsGroup(competitors)

	TestingUtils.CheckEquals(1, len(pairings), "", t)
	pair := pairings[0]

	if (pair.competitor1.name != "Benni") || (pair.competitor2.name != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)", pair.competitor1.name, pair.competitor2.name)
	}
}

func TestCalcPairings3Competitors(t *testing.T) {
	benni := Competitor{"Benni"}
	zoe := Competitor{"Zoé"}
	dani := Competitor{"Dani"}
	competitors := Competitors{}
	competitors.items = append(competitors.items, benni)
	competitors.items = append(competitors.items, dani)
	competitors.items = append(competitors.items, zoe)

	pairings := calcPairingsGroup(competitors)

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
		Competitor{"Benni"},
		Competitor{"Dani"},
		Competitor{"Zoé"},
		Competitor{"Mona"},
		Competitor{"Andrea"})

	pairings := calcPairingsGroup(competitors)
	fmt.Println(pairings)

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
	competitors.items = append(competitors.items,
		Competitor{"1"},
		Competitor{"2"},
		Competitor{"3"},
		Competitor{"4"},
		Competitor{"5"},
		Competitor{"6"},
		Competitor{"7"},
		Competitor{"8"},
		Competitor{"9"},
		Competitor{"10"},
		Competitor{"11"},
		Competitor{"12"},
	)

	pairings := calcPairingsGroup(competitors)
	fmt.Println(pairings)
	TestingUtils.CheckEquals(66, len(pairings), "", t)
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same Competitors found! " + msg)
	}
}
