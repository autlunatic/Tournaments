package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcPairings(t *testing.T) {
	benni := Contributor{"Benni"}
	dani := Contributor{"Dani"}
	contributors := Contributors{}
	contributors.items = append(contributors.items, benni)
	contributors.items = append(contributors.items, dani)

	pairings := calcPairingsGroup(contributors)

	TestingUtils.CheckEquals(1, len(pairings), "", t)
	pair := pairings[0]

	if (pair.contributor1.name != "Benni") || (pair.contributor2.name != "Dani") {
		t.Errorf("first contriburor should be Benni (but was %s) and second should be Dani (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
}

func TestCalcPairings3Contributors(t *testing.T) {
	benni := Contributor{"Benni"}
	zoe := Contributor{"Zoé"}
	dani := Contributor{"Dani"}
	contributors := Contributors{}
	contributors.items = append(contributors.items, benni)
	contributors.items = append(contributors.items, dani)
	contributors.items = append(contributors.items, zoe)

	pairings := calcPairingsGroup(contributors)

	TestingUtils.CheckEquals(3, len(pairings), "", t)
	pair := pairings[0]
	if (pair.contributor1.name != "Benni") || (pair.contributor2.name != "Dani") {
		t.Errorf("first contriburor should be Benni (but was %s) and second should be Dani (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
	pair = pairings[1]
	if (pair.contributor1.name != "Benni") || (pair.contributor2.name != "Zoé") {
		t.Errorf("first contriburor should be Benni (but was %s) and second should be Zoé (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
	pair = pairings[2]
	if (pair.contributor1.name != "Dani") || (pair.contributor2.name != "Zoé") {
		t.Errorf("first contriburor should be Dani (but was %s) and second should be Zoé (but was %s)", pair.contributor1.name, pair.contributor2.name)
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
			return pair.contributor1.name + " vs. " + pair.contributor2.name
		}

	}
	return ""
}
func isSamePair(p pairing, p2 pairing) bool {
	return (p.contributor1.name == p2.contributor1.name) && (p.contributor2.name == p2.contributor2.name) ||
		(p.contributor1.name == p2.contributor2.name) && (p.contributor2.name == p2.contributor1.name)

}

func TestCalcPairings5Contributors(t *testing.T) {
	contributors := Contributors{}
	contributors.items = append(contributors.items,
		Contributor{"Benni"},
		Contributor{"Dani"},
		Contributor{"Zoé"},
		Contributor{"Mona"},
		Contributor{"Andrea"})

	pairings := calcPairingsGroup(contributors)

	TestingUtils.CheckEquals(10, len(pairings), "", t)
	pair := pairings[0]
	if (pair.contributor1.name != "Benni") || (pair.contributor2.name != "Dani") {
		t.Errorf("first contriburor should be Benni (but was %s) and second should be Dani (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
	pair = pairings[1]
	if (pair.contributor1.name != "Mona") || (pair.contributor2.name != "Andrea") {
		t.Errorf("first contriburor should be Mona (but was %s) and second should be Andrea (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
	pair = pairings[2]
	if (pair.contributor1.name != "Benni") || (pair.contributor2.name != "Zoé") {
		t.Errorf("first contriburor should be Benni (but was %s) and second should be Zoé (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same Contributors found! " + msg)
	}
}
func TestCalcPairings12Contributors(t *testing.T) {
	contributors := Contributors{}
	contributors.items = append(contributors.items,
		Contributor{"1"},
		Contributor{"2"},
		Contributor{"3"},
		Contributor{"4"},
		Contributor{"5"},
		Contributor{"6"},
		Contributor{"7"},
		Contributor{"8"},
		Contributor{"9"},
		Contributor{"10"},
		Contributor{"11"},
		Contributor{"12"},
	)

	pairings := calcPairingsGroup(contributors)

	TestingUtils.CheckEquals(66, len(pairings), "", t)
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same Contributors found! " + msg)
	}
}
