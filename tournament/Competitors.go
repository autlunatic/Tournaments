package tournament

type Competitors struct {
	items []Competitor
}

func (c Competitors) getCompetitors()[]Competitor{
	return c.items
}

type Competitor struct {
	name string
}
