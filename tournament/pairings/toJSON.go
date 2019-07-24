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
func filterActual(ps []P, det detail.D) []P {
	var out []P
	for _, p := range ps {
		start := p.StartTime.Add(-(time.Second * time.Duration(p.StartTime.Second())))
		if start.Before(time.Now()) &&
			(start.Add(time.Minute * time.Duration(det.MinutesPerGame)).After(time.Now())) {
			out = append(out, p)
		}
	}
	return out
}

func filterOld(ps []P, det detail.D) []P {
	var out []P
	for _, p := range ps {
		gameDuration := time.Minute * time.Duration(det.MinutesPerGame)
		now := time.Now().Add(-gameDuration)
		start := p.StartTime.Add(-(time.Second * time.Duration(p.StartTime.Second())))
		if start.Before(now) &&
			(start.Add(gameDuration).After(now)) {
			out = append(out, p)
		}
	}
	return out
}

// FilterActualPairings returns the pairings that are played now
func FilterActualPairings(c []competitors.C, ap []P, finals []P, det detail.D) []P {
	ap = append(ap, finals...)
	actualPairs := filterActual(ap, det)
	return actualPairs
}

// FilterOldPairings returns the pairings that where played just before the actual round
func FilterOldPairings(c []competitors.C, ap []P, finals []P, det detail.D) []P {
	ap = append(ap, finals...)
	actualPairs := filterOld(ap, det)
	return actualPairs
}
