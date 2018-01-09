package groups

import (
	"math"

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

// CalcFinalistsPairingIDs returns an array of int where the Position on the Finaltournament plan should be
// this plan is optimized that the best players are not playing against each other as long as possible
func CalcFinalistsPairingIDs(finalRounds int) []int {
	out := make([]int, int(math.Pow(2, float64(finalRounds))))
	out[0] = 1
	out[1] = 2
	for fr := 1; fr < finalRounds; fr++ {
		// Mirror Array
		halfSlice := int(math.Pow(2, float64(fr)))
		for i := 0; i < halfSlice; i++ {
			out[i+halfSlice] = out[halfSlice-i-1]
		}
		// L x -> L x*2
		for i := 0; i < int(math.Pow(2, float64(fr+1))); i++ {
			out[i] = out[i] * 2
			if (i+1)%2 == 1 {
				out[i]--
			}
		}
	}
	return out
}
