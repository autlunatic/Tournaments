package tournament

import "strconv"

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
	i, err := strconv.Atoi(name)
	if err == nil {
		c.drawNumber = i
	} else {
		c.drawNumber = 0
	}
	return *c
}

type Competitor struct {
	name        string
	groupPoints int
	drawNumber  int
}
