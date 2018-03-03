package pairings

import (
	"time"

	"github.com/autlunatic/Tournaments/tournament/detail"
)

// CalcTimesForFinalPairings sets the time for the finalpairings and returns them again
// if FinalsStartTime in detail.D is set this will be used, if not the lastGameGroup will be used
// this function starts a new round when no courts are available and when a new finalround is started
func CalcTimesForFinalPairings(lastGameGroup time.Time, finPairs []P, d detail.D) []P {
	start := lastGameGroup
	duration := time.Minute * time.Duration(d.MinutesPerGame)
	if !d.FinalsStartTime.IsZero() {
		start = d.FinalsStartTime
	} else {
		start = start.Add(duration)
	}
	var out []P

	var count int
	var lastRound int
	for _, p := range finPairs {
		if count >= d.NumberOfParallelGames ||
			(lastRound != p.Round && lastRound != 0) {
			start = start.Add(duration)
			count = 0
		}
		count++
		lastRound = p.Round
		p.StartTime = start

		out = append(out, p)
	}
	return out
}
