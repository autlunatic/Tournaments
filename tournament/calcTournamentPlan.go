package tournament

import (
	"sort"
)

type sortByRound []pairing

func (a sortByRound) Len() int           { return len(a) }
func (a sortByRound) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByRound) Less(i, j int) bool { return a[i].round < a[j].round }

func calcTournamentPlan(groups []group, countOfParallelGames int) [][]pairing {

	var allPairs sortByRound
	for _, g := range groups {
		pairs := g.getPairings()
		for _, p := range pairs {
			allPairs = append(allPairs, p)
		}
	}

	sort.Sort(allPairs)
	var round []pairing
	var result [][]pairing
	i := 0

	for _, p := range allPairs {
		if len(round) >= countOfParallelGames ||
			foundCompetitorInRound(p, round) ||
			pairingShouldBePlayedInNextRoundForAllOfGroupSimultaneously(groups, round, p, countOfParallelGames) {
			result = append(result, round)
			i++
			round = []pairing{}
		}
		round = append(round, p)
	}
	if len(round) >= 1 {
		result = append(result, round)
	}
	return result
}
func pairingShouldBePlayedInNextRoundForAllOfGroupSimultaneously(groups []group, round []pairing, ap pairing, countOfParallelGames int) bool {
	// if any pairing of the group is already in the round then this one can be played too
	for _, g := range groups {
		gPairings := g.getPairings()
		if isPairingInPairings(ap, gPairings) {
			// group found
			if !g.isLastRound(ap) {
				return false
			}

			for _, p := range gPairings {
				if isPairingInPairings(p, round) {
					return false
				}
			}

			gamesperRound := g.getGamesPerRound()
			if len(round)+gamesperRound <= countOfParallelGames {
				return false
			}

		}
	}
	// no one of the group is already in round, check if all of the groupround fits in the free fields

	return true
}

func isPairingInPairings(p pairing, ps []pairing) bool {
	for _, mp := range ps {
		if mp.Equals(p) {
			return true
		}
	}
	return false
}

func foundCompetitorInRound(p pairing, round []pairing) bool {
	for _, r := range round {
		if (r.competitor1 == p.competitor1) ||
			(r.competitor1 == p.competitor2) ||
			(r.competitor2 == p.competitor1) ||
			(r.competitor2 == p.competitor2) {
			return true
		}
	}
	return false
}
