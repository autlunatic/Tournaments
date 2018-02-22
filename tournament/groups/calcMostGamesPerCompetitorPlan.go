package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// CalcMostGamesPerCompetitorPlan combines all of the functions and returns the complete info for the tournament Groupphase
func CalcMostGamesPerCompetitorPlan(cg []competitors.C, details detail.D) ([][]pairings.P, []G, []pairings.P) {
	var plan [][]pairings.P
	var allPairs []pairings.P
	var groups []G
	var err error
	for i := 1; ; i++ {
		groups, err = CalcGroups(cg, i)
		if err == nil {
			plan, allPairs = CalcPlan(cg, groups, details)
			if len(plan)*details.MinutesPerGame <= details.MinutesAvailForGroupsPhase {
				break
			}
		} else {
			return nil, nil, nil
		}
	}
	return plan, groups, allPairs
}
