package pairings

import (
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

// RecalcFinals calculates all the final rounds it stops when no results are available
func RecalcFinals(pairs []P, res Results, calcer tournamentPoints.TournamentPointCalcer) []P {
	calcedPairs := filterOutFirstFinalRound(pairs)
	out := calcedPairs
	for {
		calcedPairs = CalcNextFinalRound(calcedPairs, res, calcer)
		if len(calcedPairs) == 0 {
			return doSortByIDDesc(out)
		}
		out = append(out, calcedPairs...)
	}
}

// CalcNextFinalRound creates a slice of pairings for the next final round
// e.g. from quarterfinals to semifinals
// it also can be used if no all Results are given to calculate the plan before the finals are over
func CalcNextFinalRound(pairs []P, res Results, calcer tournamentPoints.TournamentPointCalcer) []P {
	if moreThanOneFinalRound(pairs) {
		return []P{}
	}
	mOut := make([]P, len(pairs)/2)
	for i := range mOut {
		mOut[i].Competitor1ID = -1
		mOut[i].Competitor2ID = -1
	}

	var oneResFound bool
	minPairID, maxPairID := calcMinMaxPairID(pairs)
	for _, p := range pairs {
		if r, ok := res[p.ID]; ok {
			oneResFound = true
			oid := (-p.ID + maxPairID) / 2
			mOut[oid].ID = minPairID - (oid) - 1
			mOut[oid].Round = -len(pairs) / 2
			r1, r2 := calcer.Calc(r.GamePoints1, r.GamePoints2)
			calcAndSetCompetitorIds(&mOut[oid], r1, r2, p)
		}
	}
	if !oneResFound {
		return []P{}
	}
	var out []P
	for _, p := range mOut {
		if p.Round < 0 {
			out = append(out, p)
		}
	}
	return out
}

func moreThanOneFinalRound(pairs []P) bool {
	oldRound := -999999
	for _, p := range pairs {
		if oldRound != -999999 && oldRound != p.Round {
			return true
		}
		oldRound = p.Round
	}
	return false
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

func filterOutFirstFinalRound(ps []P) []P {
	ps = doSortByIDDesc(ps)
	var out []P
	var lastRound int
	for i, p := range ps {
		if lastRound < p.Round {
			return out
		}
		lastRound = p.Round
		out = append(out, ps[i])
	}
	return out
}
