package groups

import (
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
	//	for gi := range c.groups
}

// CalcPairingsForFinals generates the pairings for the finalists, it takes in account that no one should play
// against an competitor wich he already faced in groupphase
func CalcPairingsForFinals(groups []G, finalistCount int) []pairings.P {
	calc := newCalcForFinals(groups)
	calc.doCalc()
	return calc.out
}
