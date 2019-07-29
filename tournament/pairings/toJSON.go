package pairings

import (
	"encoding/json"
	"time"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
)

// AllPairsJSON generates the json for the complete Gameplan including all Finals
func AllPairsJSON(c []competitors.C, ap []P, finals []P, numberOfParallelGames int, groupText string) string {
	var pairSections []PairingInfoSection
	// group := CalcedPlanToGamePlan(c, plan)
	group := AllPairsToGamePlan(c, ap)
	pairSections = append(pairSections, PairingInfoSection{groupText, group.PairingInfo})
	fin := AllPairsToGamePlan(c, finals)
	pairSections = append(pairSections, PairingInfoSection{"Finalrunden", fin.PairingInfo})
	out, err := json.Marshal(pairSections)
	if err != nil {
		return ""
	}
	return string(out)

}
func filterForTimeStamp(ps []P, det detail.D, now time.Time) []P {
	var out []P
	for _, p := range ps {
		start := p.StartTime.Add(-(time.Second * time.Duration(p.StartTime.Second())))
		if start.Before(now) &&
			(start.Add(time.Minute * time.Duration(det.MinutesPerGame)).After(now)) {
			out = append(out, p)
		}
	}
	return out
}

// FilterActualPairings returns the pairings that are played now
func FilterActualPairings(c []competitors.C, ap []P, finals []P, det detail.D) []P {
	ap = append(ap, finals...)
	return filterForTimeStamp(ap, det, time.Now())
}

// FilterOldPairings returns the pairings that where played just before the actual round
func FilterOldPairings(c []competitors.C, ap []P, finals []P, det detail.D) []P {
	ap = append(ap, finals...)
	gameDuration := time.Minute * time.Duration(det.MinutesPerGame)
	return filterForTimeStamp(ap, det, time.Now().Add(-gameDuration))
}

// FilterComingPairings returns the pairings that where played just before the actual round
func FilterComingPairings(c []competitors.C, ap []P, finals []P, det detail.D) []P {
	ap = append(ap, finals...)
	gameDuration := time.Minute * time.Duration(det.MinutesPerGame)
	return filterForTimeStamp(ap, det, time.Now().Add(gameDuration))
}
