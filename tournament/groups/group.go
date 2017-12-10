package groups

import (
	"github.com/autlunatic/tournaments/tournament/competitors"
	"github.com/autlunatic/tournaments/tournament/pairings"
)

// Group represents a Group of a Tournament in group phase
// it holds a id which should be unique and the competitors in this group
type Group struct {
	id          int
	competitors competitors.Competitors
}

func (g Group) getPairings() []pairings.Pairing {
	return pairings.CalcPairings(g.competitors.Items, g.id)
}

func (g Group) isLastRound(ap pairings.Pairing) bool {
	p := g.getPairings()
	return pairings.GetMaxRoundOfPairings(p) == ap.Round
}

func (g Group) getGamesPerRound() int {
	ps := g.getPairings()
	m := make(map[int]int)
	for _, p := range ps {
		m[p.Round]++
	}
	return m[1]
}
