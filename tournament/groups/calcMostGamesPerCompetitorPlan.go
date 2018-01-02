package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

func calcMostGamesPerCompetitorPlan(cg []competitors.C, details detail.Details) ([][]pairings.P, []G) {
	var plan [][]pairings.P
	var groups []G
	var err error
	for i := 1; ; i++ {
		groups, err = CalcGroups(cg, i)
		if err == nil {
			plan = calcPlan(cg, groups, details.NumberOfParallelGames)
			if len(plan)*details.MinutesPerGame <= details.MinutesAvailForGroupsPhase {
				break
			}
		} else {
			return nil, nil
		}
	}
	return plan, groups
}
