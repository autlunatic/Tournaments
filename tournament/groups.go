package tournament

type group struct {
	id          int
	competitors Competitors
}

func (g group) getPairings() []pairing {
	return calcPairings(g.competitors.items)
}

func (g group) isLastRound(ap pairing) bool {
	pairings := g.getPairings()
	maxRound := 0
	for _, p := range pairings {
		if maxRound < p.round {
			maxRound = p.round
		}
	}
	return maxRound == ap.round
}

func (g group) getGamesPerRound() int {
	ps := g.getPairings()
	m := make(map[int]int)
	for _, p := range ps {
		m[p.round]++
	}
	return m[1]
}
