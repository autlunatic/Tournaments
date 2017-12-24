package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

func calcMostGamesPerCompetitorPlan(cg []competitors.Competitor, details detail.Details) ([][]pairings.Pairing, []Group) {
	var plan [][]pairings.Pairing
	var groups []Group
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
