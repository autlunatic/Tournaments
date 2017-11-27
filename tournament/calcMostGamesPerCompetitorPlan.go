package tournament

func calcMostGamesPerCompetitorPlan(cg CompetitorsGetter, details Details) ([][]pairing, []group) {
	var plan [][]pairing
	var groups []group
	var err error
	for i := 1; ; i++ {
		groups, err = calcGroups(cg, i)
		if err == nil {
			plan = calcPlan(groups, details.numberOfParallelGames)
			if len(plan) * details.minutesPerGame <= details.minutesAvailForGroupsPhase {
				break
			}
		} else {
			return nil, nil
		}
	}
	return plan, groups
}
