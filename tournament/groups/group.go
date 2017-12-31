package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// Group represents a Group of a Tournament in group phase
// it holds a id which should be unique and the competitors in this group
type Group struct {
	id          int
	Competitors []competitors.Competitor
}

// AddCompetitors adds the given competitors to the Group
func (g *Group) AddCompetitors(comps []competitors.Competitor) {
	for i := range comps {
		g.Competitors = append(g.Competitors, comps[i])
	}
}

func (g Group) getPairings() ([]pairings.Pairing, error) {
	return pairings.CalcPairings(g.Competitors, g.id)
}

func (g Group) isLastRound(ap pairings.Pairing) bool {
	p, err := g.getPairings()
	if err != nil {
		return true
	}
	return pairings.GetMaxRoundOfPairings(p) == ap.Round
}

func (g Group) getGamesPerRound() int {
	ps, err := g.getPairings()
	if err != nil {
		return 0
	}
	m := make(map[int]int)
	for _, p := range ps {
		m[p.Round]++
	}
	return m[1]
}
