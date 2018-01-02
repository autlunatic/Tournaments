package groups

import (
	"sort"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

type planCalc struct {
	groups               []G
	allCompetitors       []competitors.C
	countOfParallelGames int
}

func calcPlan(allCompetitors []competitors.C, Groups []G, countOfParallelGames int) [][]pairings.P {

	calc := planCalc{
		groups:               Groups,
		countOfParallelGames: countOfParallelGames,
		allCompetitors:       allCompetitors,
	}

	allPairs, err := calc.calcPairsFromGroups()
	if err != nil {
		return nil
	}

	for i := range allPairs {
		allPairs[i].ID = i
	}

	var round []pairings.P
	var result [][]pairings.P
	i := 0

	for _, p := range allPairs {
		if calc.needNewGroup(round, p) {
			result = append(result, round)
			i++
			round = []pairings.P{}
		}
		round = append(round, p)
	}
	if len(round) >= 1 {
		result = append(result, round)
	}
	return result
}

func (c planCalc) needNewGroup(round []pairings.P, p pairings.P) bool {
	roundFull := len(round) >= c.countOfParallelGames
	foundSame := foundSameCompetitorInRound(p, round)
	lastRoundSimultaneously := c.pairingShouldBePlayedInNextRoundForAllOfGroupSimultaneously(round, p)
	return roundFull || foundSame || lastRoundSimultaneously
}

func (c planCalc) calcPairsFromGroups() ([]pairings.P, error) {
	var allPairs pairings.SortByRoundGroupDrawnumber
	allPairs.Comps = c.allCompetitors
	for _, g := range c.groups {
		pairs, err := g.getPairings()
		if err != nil {
			return nil, err
		}
		allPairs.Pairs = append(allPairs.Pairs, pairs...)
	}
	sort.Sort(allPairs)

	return allPairs.Pairs, nil
}

func (c planCalc) pairingShouldBePlayedInNextRoundForAllOfGroupSimultaneously(round []pairings.P, ap pairings.P) bool {
	// if any pairing of the group is already in the round then this one can be played too
	for _, g := range c.groups {
		gPairings, _ := g.getPairings()
		if ap.InPairings(gPairings) {
			// Group found
			if !g.isLastRound(ap) {
				return false
			}
			// if in last round and other of this Group are already in round no new Group is needed
			for _, p := range gPairings {
				if p.InPairings(round) {
					return false
				}
			}
			// no one of the Group is already in round, check if all of the groupRound fits in the free fields
			gamesPerRound := g.getGamesPerRound()
			if len(round)+gamesPerRound <= c.countOfParallelGames {
				return false
			}
			break
		}
	}
	return true
}

func foundSameCompetitorInRound(p pairings.P, round []pairings.P) bool {
	for _, r := range round {
		if (r.Competitor1ID == p.Competitor1ID) ||
			(r.Competitor1ID == p.Competitor2ID) ||
			(r.Competitor2ID == p.Competitor1ID) ||
			(r.Competitor2ID == p.Competitor2ID) {
			return true
		}
	}
	return false
}
