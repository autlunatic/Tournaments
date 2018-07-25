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

// ActualPairsJSON returns the pairings that are played now
func ActualPairsJSON(c []competitors.C, ap []P, finals []P, det detail.D) string {
	ap = append(ap, finals...)
	group := filterActual(ap, det)
	return AllPairsJSON(c, group, []P{}, det.NumberOfParallelGames, "")
}
