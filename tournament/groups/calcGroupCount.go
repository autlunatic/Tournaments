package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
)

// CalcGroupCountMaxGames calculates the possible count of groups given in that time slot.
// this func focuses on having as much games as possible by making the count of competitors per Group as big as possible
func CalcGroupCountMaxGames(c []competitors.C, totalTimeMinutes int, minutesPerGame int, countOfParallelGames int) int {
	_, g, _ := CalcMostGamesPerCompetitorPlan(c, *detail.New(countOfParallelGames, minutesPerGame, totalTimeMinutes, 0))
	return len(g)
}
