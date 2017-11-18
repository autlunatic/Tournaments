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

	pairingCount := calcPairingsGroup(contributors)

	TestingUtils.CheckEquals(2, len(pairingCount), t)

}
