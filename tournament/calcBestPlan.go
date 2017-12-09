package tournament



func calcBestPlan(cg CompetitorsGetter, details Details) ([][]pairing,[]group){
	var plan [][]pairing
	var groups []group
	for i:=1;;i++{
		groups = calcGroups(cg, i)
		plan = calcPlan(groups, details.numberOfParallelGames)
		duration := len(plan) * details.minutesPerGame
		if duration <= details.minutesAvailForGroupsPhase{
			break
		}
	}
	return plan, groups
}
