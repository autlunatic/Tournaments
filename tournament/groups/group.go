package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// G represents a Group of a Tournament in group phase
// it holds a id which should be unique and the competitors in this group
type G struct {
	id          int
	Competitors []competitors.C
}

// AddCompetitors adds the given competitors to the Group
func (g *G) AddCompetitors(comps []competitors.C) {
	for i := range comps {
		g.Competitors = append(g.Competitors, comps[i])
	}
}

func (g G) getPairings() ([]pairings.P, error) {
	return pairings.CalcPairings(g.Competitors, g.id)
}

func (g G) isLastRound(ap pairings.P) bool {
	p, err := g.getPairings()
	if err != nil {
		return true
	}
	return pairings.GetMaxRoundOfPairings(p) == ap.Round
}

func (g G) getGamesPerRound() int {
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
