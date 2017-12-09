package tournament


//Competitors hold a slice of Competitor
type Competitors struct {
	items []Competitor
}

// CompetitorsGetter is used when we need a slice of competitors
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
// Competitor holds the data for one team/person
type Competitor struct {
	name        string
	groupPoints int
	drawNumber  int
}
