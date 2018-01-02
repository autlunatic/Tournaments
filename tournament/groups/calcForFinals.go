package groups

import (
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

type calcForFinals struct {
	groups []G
	out    []pairings.P
}

func newCalcForFinals(groups []G) *calcForFinals {
	out := new(calcForFinals)
	out.groups = groups
	return out
}

func (c *calcForFinals) doCalc() {
	//	for gi := range c.groups
}

// CalcForFinals generates the pairings for the finalists, it takes in account that no one should play
// against an competitor wich he already faced in groupphase
func CalcForFinals(groups []G) []pairings.P {
	calc := newCalcForFinals(groups)

	calc.doCalc()
	return calc.out

}
