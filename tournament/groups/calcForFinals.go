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
	fin := DetermineFinalists(c.groups, c.finalistCount)
	cs := competitors.GetCompetitorsSortedByPlacementAndGroupPoints(fin)
	c.out = make([]pairings.P, c.finalistCount/2)
	finalistsPairingIDs := CalcFinalistsPairingIDs(int(math.Log2(float64(c.finalistCount))))
	for i := 0; i < c.finalistCount; i++ {
		c.setPairingForSortedID(finalistsPairingIDs[i], cs[i].ID())
	}
}

func (c *calcPairingsForFinals) setPairingForSortedID(sID int, compID int) {
	pairingID := (sID - 1) / 2
	competID := sID % 2
	if competID == 1 {
		c.out[pairingID].Round = -c.finalistCount
		c.out[pairingID].ID = -pairingID - 1
		c.out[pairingID].Competitor1ID = compID
	} else {
		c.out[pairingID].Competitor2ID = compID
	}
}

// CalcPairingsForFinals generates the pairings for the finalists, it takes in account that no one should play
// against an competitor wich he already faced in groupphase
func CalcPairingsForFinals(groups []G, finalistCount int) ([]pairings.P, error) {
	calc := newCalcForFinals(groups)
	calc.finalistCount = finalistCount
	calc.doCalc()
	return calc.out, nil
}

// CalcFinalistsPairingIDs returns an slice of int where the Position on the Finaltournament plan should be
// this plan is optimized that the best players are not playing against each other as long as possible
func CalcFinalistsPairingIDs(finalRounds int) []int {
	out := make([]int, int(math.Pow(2, float64(finalRounds))))
	out[0] = 1
	out[1] = 2
	for fr := 1; fr < finalRounds; fr++ {
		// Mirror Slice
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
