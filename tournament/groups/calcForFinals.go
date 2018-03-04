package groups

import (
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// CalcPairingsForFinals generates the pairings for the finalists, it takes in account that no one should play
// against an competitor wich he already faced in groupphase
func CalcPairingsForFinals(g []G, finalistCount int) ([]pairings.P, error) {
	if getCompetitorCount(g) < finalistCount {
		return nil, fmt.Errorf("error! more finalists than competitors")
	}
	calc := newCalcForFinals(g)
	calc.finalistCount = finalistCount
	calc.doCalc()

	return calc.out, nil
}

// CalcFinalistRankings returns an slice of int where the Position on the Finaltournament plan should be
// this plan is optimized that the best players are not playing against each other as long as possible
func CalcFinalistRankings(finalRounds int) []int {
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

type calcPairingsForFinals struct {
	groups        []G
	finalistCount int
	out           []pairings.P
}

func newCalcForFinals(groups []G) *calcPairingsForFinals {
	out := &calcPairingsForFinals{groups: groups}
	return out
}

func isInSlice(slice []int, val int) bool {
	for _, i := range slice {
		if i == val {
			return true
		}
	}
	return false
}

func (c *calcPairingsForFinals) doCalc() {
	fin := DetermineFinalists(c.groups, c.finalistCount)
	cs := competitors.SortedByPlacementAndPoints(fin)
	c.out = make([]pairings.P, c.finalistCount/2)
	finRankPos := CalcFinalistRankings(int(math.Log2(float64(c.finalistCount))))
	// first fill the better half, the better half should be fixed,
	// the lower half should be mixed if there would be a pairing of two competitors of the same group
	for i := 0; i < c.finalistCount/2; i++ {
		err := c.setPairingForSortedID(finRankPos[i], cs[i].ID(), false)
		if err != nil {
			log.Println("Error occurred while filling first half")
		}
	}
	err := c.setLowerHalf(finRankPos, cs, false)
	if err != nil {
		err := c.setLowerHalf(finRankPos, cs, true)
		if err != nil {
			log.Println("Error occurred while filling lower half reversed")
		}
	}
}

func (c *calcPairingsForFinals) calcIndex(reversed bool, index int) int {
	var out int
	if reversed {
		out = c.finalistCount - 1 - index + c.finalistCount/2
	} else {
		out = index
	}
	return out

}
func (c *calcPairingsForFinals) setLowerHalf(finRankPos []int, cs []competitors.C, fillReversed bool) error {
	var usedRanks []int
	for i := c.finalistCount - 1; i >= c.finalistCount/2; i-- {
		c.setNextRankingIndexForCompetitor(finRankPos, cs, fillReversed, i, &usedRanks)
	}
	if len(usedRanks) != len(finRankPos)/2 {
		return errors.New("there are competitors playing against others of the same group")
	}
	return nil
}

func (c *calcPairingsForFinals) setNextRankingIndexForCompetitor(finRankPos []int, cs []competitors.C, fillReversed bool, i int, usedRanks *[]int) {
	for r := c.finalistCount - 1; r >= c.finalistCount/2; r-- {
		ri := c.calcIndex(fillReversed, r)
		ii := c.calcIndex(fillReversed, i)
		if isInSlice(*usedRanks, finRankPos[ri]) {
			continue
		}
		err := c.setPairingForSortedID(finRankPos[ri], cs[ii].ID(), true)
		if err == nil {
			*usedRanks = append(*usedRanks, finRankPos[ri])
			break
		}
	}
}

func (c *calcPairingsForFinals) setPairingForSortedID(sID int, compID int, checkForDuplicates bool) error {
	pID := (sID - 1) / 2
	competID := sID % 2
	c.out[pID].Round = -c.finalistCount / 2
	c.out[pID].ID = -pID - 1
	if competID == 1 {
		c.out[pID].Competitor1ID = compID
	} else {
		c.out[pID].Competitor2ID = compID
	}
	if checkForDuplicates {
		err := c.doCheckForDuplicates(pID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *calcPairingsForFinals) doCheckForDuplicates(pID int) error {
	g1, err1 := GetGroupIDOfCompetitor(c.groups, c.out[pID].Competitor1ID)
	if err1 != nil {
		return err1
	}
	g2, err2 := GetGroupIDOfCompetitor(c.groups, c.out[pID].Competitor2ID)
	if err2 != nil {
		return err2
	}
	if len(c.groups) > 1 && g1 == g2 {
		return fmt.Errorf("Same Group not possible C1: %v C2: %v both in group: %v", c.out[pID].Competitor1ID, c.out[pID].Competitor2ID, g1)
	}
	return nil
}
