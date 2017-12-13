package pairings

import (
	"errors"
	"sort"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

type calcGroupPlacements struct {
	pairings      []Pairing
	pairingResult []PairingResult
}

// Placement holds information about the result placement
type Placement struct {
	competitor competitors.Competitor
	placement  int
}

// NewPlacement creates and initializes a placement
func NewPlacement(c competitors.Competitor, place int) Placement {
	var out Placement
	out.competitor = c
	out.placement = place
	return out
}

// SortByPoints implements the interface interface for a Placement slice
type SortByPoints []Placement

func (s SortByPoints) Len() int {
	return len(s)
}

func (s SortByPoints) Less(i, j int) bool {
	return s[i].competitor.GroupPoints > s[j].competitor.GroupPoints
}
func (s SortByPoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (cgp calcGroupPlacements) calcPlacements(tpc tournamentPoints.TournamentPointCalcer) ([]Placement, error) {
	var out []Placement
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
	for k, v := range compMap {
		k.GroupPoints = v
		out = append(out, NewPlacement(k, 0))
	}
	sort.Sort(SortByPoints(out))
	for i := range out {
		out[i].placement = i + 1
	}
	return out, nil
}
