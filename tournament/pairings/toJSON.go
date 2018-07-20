package pairings

import (
	"encoding/json"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// AllPairsJSON generates the json for the complete Gameplan including all Finals
func AllPairsJSON(c []competitors.C, ap []P, finals []P, numberOfParallelGames int) string {
	var pairSections []PairingInfoSection
	fin := AllPairsToGamePlan(c, finals)
	// group := CalcedPlanToGamePlan(c, plan)
	group := AllPairsToGamePlan(c, ap)
	pairSections = append(pairSections, PairingInfoSection{"Gruppenphase", group.PairingInfo})
	pairSections = append(pairSections, PairingInfoSection{"Finalrunden", fin.PairingInfo})
	out, err := json.Marshal(pairSections)
	if err != nil {
		return ""
	}
	return string(out)

}
