package pairings

import (
	"encoding/json"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// AllPairsJSON generates the json for the complete Gameplan including all Finals
func AllPairsJSON(c []competitors.C, plan [][]P, finals []P, numberOfParallelGames int) string {
	var pairSections []PairingInfoSection
	fin := AllPairsToGamePlan(c, finals, numberOfParallelGames)
	group := CalcedPlanToGamePlan(c, plan)
	pairSections = append(pairSections, PairingInfoSection{"group", group.PairingInfo})
	pairSections = append(pairSections, PairingInfoSection{"finals", fin.PairingInfo})
	out, err := json.Marshal(pairSections)
	if err != nil {
		return ""
	}
	return string(out)

}
