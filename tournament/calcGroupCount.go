package tournament

// calcGroupCountMaxGames calculates the possible count of groups given in that time slot.
// this func focuses on having as much games as possible by making the count of competitors per group as big as possible

func calcGroupCountMaxGames(c CompetitorsGetter, totalTimeMinutes int, minutesPerGame int, countOfParallelGames int) int {
	_, g := calcMostGamesPerCompetitorPlan(c, details{countOfParallelGames, minutesPerGame, totalTimeMinutes})

	return len(g)

}
