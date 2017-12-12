package pairings

import (
	"errors"

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

func (p placement) isInSlice([]placement) {

}

func (cgp calcGroupPlacements) calcPlacements(tpc tournamentPoints.TournamentPointCalcer) ([]placement, error) {
	var out []placement
	// creating a map for IDS
	prm := make(map[int]PairingResult)
	for _, pr := range cgp.pairingResult {
		prm[pr.ID] = pr
	}

	compMap := make(map[competitors.Competitor]struct{})

	for _, p := range cgp.pairings {
		_ = compMap[p.Competitor1]
		_ = compMap[p.Competitor2]

		pairingResult, ok := prm[p.ID]
		if !ok {
			return nil, errors.New("result IDs don't match the Pairings IDs, can't calc the result")
		}
		tp1, tp2 := tpc.Calc(pairingResult.gamePoints1, pairingResult.gamePoints2)
		p.Competitor1.GroupPoints += tp1
		p.Competitor2.GroupPoints += tp2
	}

	for k, _ := range compMap {
		out = append(out, NewPlacement(k, 1))
	}

	return out, nil
}
