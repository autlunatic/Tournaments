package pairings

import (
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

// CalcNextFinalRound creates a slice of pairings for the next final round
// e.g. from quarterfinals to semifinals
// it also can be used if no all Results are given to calculate the plan before the finals are over
func CalcNextFinalRound(pairs []P, res Results, calcer tournamentPoints.TournamentPointCalcer) []P {
	out := make([]P, len(pairs)/2)
	minPairID, maxPairID := calcMinMaxPairID(pairs)
	for _, p := range pairs {
		if r, ok := res[p.ID]; ok {
			oid := (-p.ID + maxPairID) / 2
			out[oid].ID = minPairID - (oid) - 1
			out[oid].Round = -len(pairs) / 2
			r1, r2 := calcer.Calc(r.gamePoints1, r.gamePoints2)
			calcAndSetCompetitorIds(&out[oid], r1, r2, p)
		}
	}
	return out
}

func calcMinMaxPairID(pairs []P) (min int, max int) {
	maxPairID := -999999
	var minPairID int
	for _, p := range pairs {
		if maxPairID < p.ID {
			maxPairID = p.ID
		}
		if minPairID > p.ID {
			minPairID = p.ID
		}
	}
	return minPairID, maxPairID
}

func calcAndSetCompetitorIds(aOut *P, r1 int, r2 int, p P) {
	// odd id -> C1 should be set
	if p.ID%2 != 0 {
		if r1 > r2 {
			aOut.Competitor1ID = p.Competitor1ID
		} else {
			aOut.Competitor1ID = p.Competitor2ID
		}
	} else {
		if r1 > r2 {
			aOut.Competitor2ID = p.Competitor1ID
		} else {
			aOut.Competitor2ID = p.Competitor2ID
		}
	}
}
