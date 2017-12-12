package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournament"
)

// calcGroupCountMaxGames calculates the possible count of groups given in that time slot.
// this func focuses on having as much games as possible by making the count of competitors per Group as big as possible

func calcGroupCountMaxGames(c competitors.Getter, totalTimeMinutes int, minutesPerGame int, countOfParallelGames int) int {
	_, g := calcMostGamesPerCompetitorPlan(c, *tournament.NewTournamentDetails(countOfParallelGames, minutesPerGame, totalTimeMinutes))

	return len(g)

}
