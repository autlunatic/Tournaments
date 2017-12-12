package pairings

import (
	"errors"
	"fmt"
	"sort"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

type calcGroupPlacements struct {
	pairings      []Pairing
	pairingResult []PairingResult
}

type placement struct {
	competitor competitors.Competitor
	placement  int
}

func NewPlacement(c competitors.Competitor, place int) placement {
	var out placement
	out.competitor = c
	out.placement = place
	return out
}

type SortByPoints []placement

func (s SortByPoints) Len() int {
	return len(s)
}

func (s SortByPoints) Less(i, j int) bool {
	return s[i].competitor.GroupPoints > s[j].competitor.GroupPoints
}
func (s SortByPoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (cgp calcGroupPlacements) calcPlacements(tpc tournamentPoints.TournamentPointCalcer) ([]placement, error) {
	var out []placement
	// creating a map for IDS
	prm := make(map[int]PairingResult)
	for _, pr := range cgp.pairingResult {
		prm[pr.ID] = pr
	}

	compMap := make(map[competitors.Competitor]int)
	for i, p := range cgp.pairings {
		pairingResult, ok := prm[p.ID]
		if !ok {
			return nil, errors.New("result IDs don't match the Pairings IDs, can't calc the result")
		}
		tp1, tp2 := tpc.Calc(pairingResult.gamePoints1, pairingResult.gamePoints2)
		compMap[cgp.pairings[i].Competitor1] += tp1
		compMap[cgp.pairings[i].Competitor2] += tp2
	}
	fmt.Println("compmap:", compMap)
	for k, v := range compMap {
		fmt.Println(k)
		k.GroupPoints = v
		out = append(out, NewPlacement(k, 0))
	}
	sort.Sort(SortByPoints(out))
	for i := range out {
		out[i].placement = i + 1
	}
	return out, nil
}
