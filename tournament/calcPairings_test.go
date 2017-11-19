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

	TestingUtils.CheckEquals(1, len(pairings), t)
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

	TestingUtils.CheckEquals(3, len(pairings), t)
	pair := pairings[0]
	if (pair.contributor1.name != "Benni") || (pair.contributor2.name != "Zoé") {
		t.Errorf("first contriburor should be Benni (but was %s) and second should be Dani (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
	pair = pairings[1]
	if (pair.contributor1.name != "Zoé") || (pair.contributor2.name != "Dani") {
		t.Errorf("first contriburor should be Zoé (but was %s) and second should be Dani (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
	pair = pairings[2]
	if (pair.contributor1.name != "Benni") || (pair.contributor2.name != "Dani") {
		t.Errorf("first contriburor should be Benni (but was %s) and second should be Dani (but was %s)", pair.contributor1.name, pair.contributor2.name)
	}
}
