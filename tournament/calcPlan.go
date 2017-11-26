package tournament

import (
	"sort"
)

type sortByRound []pairing

func (a sortByRound) Len() int           { return len(a) }
func (a sortByRound) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByRound) Less(i, j int) bool { return a[i].round < a[j].round }

type planCalc struct {
	groups               []group
	countOfParallelGames int
}

func calcPlan(groups []group, countOfParallelGames int) [][]pairing {

	calc := planCalc{
		groups:               groups,
		countOfParallelGames: countOfParallelGames,
	}

	allPairs := calc.calcPairsFromGroups()

	var round []pairing
	var result [][]pairing
	i := 0

	for _, p := range allPairs {
		if calc.needNewGroup(round, p) {
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

func (c planCalc) needNewGroup(round []pairing, p pairing) bool {
	return len(round) >= c.countOfParallelGames ||
		foundSameCompetitorInRound(p, round) ||
		c.pairingShouldBePlayedInNextRoundForAllOfGroupSimultaneously(round, p)
}

func (c planCalc) calcPairsFromGroups() sortByRound {
	var allPairs sortByRound
	for _, g := range c.groups {
		pairs := g.getPairings()
		for _, p := range pairs {
			allPairs = append(allPairs, p)
		}
	}
	sort.Sort(allPairs)
	return allPairs
}

func (c planCalc)pairingShouldBePlayedInNextRoundForAllOfGroupSimultaneously(round []pairing, ap pairing) bool {
	// if any pairing of the group is already in the round then this one can be played too
	for _, g := range c.groups {
		gPairings := g.getPairings()
		if ap.inPairings(gPairings) {
			// group found
			if !g.isLastRound(ap) {
				return false
			}
			// if in last round and other of this group are already in round no new group is needed
			for _, p := range gPairings {
				if p.inPairings(round) {
					return false
				}
			}
			// no one of the group is already in round, check if all of the groupRound fits in the free fields
			gamesPerRound := g.getGamesPerRound()
			if len(round)+gamesPerRound <= c.countOfParallelGames {
				return false
			}
			break
		}
	}
	return true
}

func foundSameCompetitorInRound(p pairing, round []pairing) bool {
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
