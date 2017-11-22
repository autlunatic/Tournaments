package tournament

type group struct {
	Id          int
	Competitors Competitors
}

func (g *group) setId(value int) {
	g.Id = value
}
