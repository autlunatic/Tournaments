package groups

import (
	"github.com/autlunatic/tournaments/tournament/competitors"
	"github.com/autlunatic/tournaments/tournament/pairings"
	"github.com/autlunatic/tournaments/tournament/tournament"
)

func calcMostGamesPerCompetitorPlan(cg competitors.Getter, details tournament.Details) ([][]pairings.Pairing, []Group) {
	var plan [][]pairings.Pairing
	var groups []Group
	var err error
	for i := 1; ; i++ {
		groups, err = CalcGroups(cg, i)
		if err == nil {
			plan = calcPlan(groups, details.NumberOfParallelGames)
			if len(plan)*details.MinutesPerGame <= details.MinutesAvailForGroupsPhase {
				break
			}
		} else {
			return nil, nil
		}
	}
	return plan, groups
}
