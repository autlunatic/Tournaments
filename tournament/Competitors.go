package tournament

import "strconv"


//Competitors hold a slice of Competitor
type Competitors struct {
	items []Competitor
}

type CompetitorsGetter interface {
	getCompetitors() []Competitor
}

func (c Competitors) getCompetitors() []Competitor {
	return c.items
}

func newCompetitor(name string) Competitor {
	c := new(Competitor)
	c.name = name
	return *c
}

type Competitor struct {
	name        string
	groupPoints int
	drawNumber  int
}
