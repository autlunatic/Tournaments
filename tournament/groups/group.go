package groups

import (
	"errors"
	"fmt"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// G represents a Group of a Tournament in group phase
// it holds a id which should be unique and the competitors in this group
type G struct {
	ID          int
	Competitors []competitors.C
}

// NewTestGroups returns groups for testing.
func NewTestGroups(count int) []G {
	out := make([]G, count)
	for i := 1; i <= count; i++ {
		out[i-1].ID = i
	}
	return out
}

// GOfCompentitorID returns the groupt of the competitor
func GOfCompentitorID(gs []G, competitorID int) (G, error) {
	for _, g := range gs {
		for _, c := range g.Competitors {
			if c.ID() == competitorID {
				return g, nil
			}
		}
	}
	return G{}, errors.New("competitorID not found in Groups")
}

// GByID returns the group of the given groupID
func GByID(gs []G, groupID int) (*G, error) {
	for i := range gs {
		if gs[i].ID == groupID {
			return &gs[i], nil
		}
	}
	return nil, fmt.Errorf("groupID  %v not found in Groups %v", groupID, gs)
}

// GetGroupIDOfCompetitor returns the id of the group in which the competitorID was found
// returns an error if competitor id was not found in the given slice
func GetGroupIDOfCompetitor(gs []G, competitorID int) (int, error) {
	for _, g := range gs {
		for _, c := range g.Competitors {
			if c.ID() == competitorID {
				return g.ID, nil
			}
		}
	}
	return -1, errors.New("competitorID not found in Groups")
}

// AddCompetitors adds the given competitors to the Group
func (g *G) AddCompetitors(comps []competitors.C) {
	for i := range comps {
		g.AddCompetitor(comps[i])
	}
}

// AddCompetitor adds the given competitor to the Group
func (g *G) AddCompetitor(comp competitors.C) {
	g.Competitors = append(g.Competitors, comp)
}

func (g G) getPairings() ([]pairings.P, error) {
	return pairings.CalcPairings(g.Competitors, g.ID)
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

func getCompetitorCount(gs []G) int {
	var out int
	for _, g := range gs {
		out += len(g.Competitors)
	}
	return out
}
