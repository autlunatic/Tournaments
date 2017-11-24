package tournament

// calcGroupCountMaxGames calculates the possible count of groups given in that time slot.
// this func focuses on having as much games as possible by making the count of competitors per group as big as possible

func calcGroupCountMaxGames(c CompetitorsGetter, totalTimeMinutes int, minutesPerGame int, countOfParallelGames int) int {
	gc := c.getCompetitors()
	for i := 1; i <= len(gc)/2; i++ {
		gct := gc[0 : len(gc)/i]
		roundCount := calcRoundsFromPairings(calcPairings(gct))
		if roundCount*minutesPerGame <= totalTimeMinutes {
			return i
		}
	}
	return 0
}
