package tournament

import "sort"

type sortByRound []pairing

func (a sortByRound) Len() int           { return len(a) }
func (a sortByRound) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByRound) Less(i, j int) bool { return a[i].round < a[j].round }

func calcTournamentPlan(groups []group, countOfParallelGames int) [][]pairing {

	var allPairs sortByRound
	for _, g := range groups {
		pair := calcPairings(g.Competitors.items)
		for _, p := range pair {
			allPairs = append(allPairs, p)
		}
	}
	sort.Sort(allPairs)
	var round []pairing
	var result [][]pairing
	i := 0

	for _, p := range allPairs {
		if len(round) >= countOfParallelGames ||
			foundCompetitorInRound(p, round) {
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
