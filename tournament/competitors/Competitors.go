package competitors

//Competitors hold a slice of Competitor that implements the Getter
type Competitors struct {
	Items []Competitor
}

// Getter provides a interface to get a slice of Competitors
type Getter interface {
	GetCompetitors() []Competitor
}

// GetCompetitors implements Getter Interface
func (c Competitors) GetCompetitors() []Competitor {
	return c.Items
}

// NewCompetitor generates a New Competitor with the given Name
func NewCompetitor(name string) Competitor {
	c := new(Competitor)
	c.Name = name
	return *c
}

// Competitor holds information
type Competitor struct {
	Name        string
	GroupPoints int
	DrawNumber  int
}
