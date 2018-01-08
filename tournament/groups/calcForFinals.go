package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

type calcPairingsForFinals struct {
	groups        []G
	finalistCount int
	out           []pairings.P
}

func newCalcForFinals(groups []G) *calcPairingsForFinals {
	out := new(calcPairingsForFinals)
	out.groups = groups
	return out
}

func (c *calcPairingsForFinals) doCalc() {
	cs := DetermineFinalists(c.groups, c.finalistCount)
	cs = competitors.GetCompetitorsSortedByGroupPoints(cs)
	c.out = make([]pairings.P, c.finalistCount/2)

	firstPlacedCount := getFirstPlacedCount(cs)
	for i := 0; i < firstPlacedCount; i++ {

	}
}

func getFirstPlacedCount(cs []competitors.C) int {
	var out int
	for _, c := range cs {
		if c.GroupPlacement() == 1 {
			out++
		}
	}
	return out
}

// CalcPairingsForFinals generates the pairings for the finalists, it takes in account that no one should play
// against an competitor wich he already faced in groupphase
func CalcPairingsForFinals(groups []G, finalistCount int) ([]pairings.P, error) {
	calc := newCalcForFinals(groups)
	calc.doCalc()
	return calc.out, nil
}
