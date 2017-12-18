package competitors

// Items contains the Competitors of a tournament
var Items Competitors

//Competitors hold a slice of SimpleCompetitor that implements the Getter
type Competitors struct {
	Items []Competitor
}

// Getter provides a interface to get a slice of Competitors
type Getter interface {
	GetCompetitors() []Competitor
}

// GetCompetitor gets the Competitor of the Items with the given ID
func GetCompetitor(ID int) Competitor {
	return Items.Items[ID]
}

// GetCompetitors implements Getter Interface
func (c Competitors) GetCompetitors() []Competitor {
	return c.Items
}

// NewCompetitor generates a New SimpleCompetitor with the given Name
func NewCompetitor(name string, id int) *SimpleCompetitor {
	c := new(SimpleCompetitor)
	c.name = name
	c.id = id
	return c

}

// Competitor is the interface used by the Tournament package for Operations with the Competitor
type Competitor interface {
	AddPoints(p int)
	GetPoints() int
	ID() int
	Name() string
	DrawNumber() int
	SetDrawNumber(int)
}

// SimpleCompetitor holds information for an minimalistic Competitor
type SimpleCompetitor struct {
	id          int
	name        string
	GroupPoints int
	drawNumber  int
}

// DrawNumber is for implementing the Competitor interface
func (c *SimpleCompetitor) DrawNumber() int {
	return c.drawNumber
}

// Name is for implementing the Competitor Interface
func (c *SimpleCompetitor) Name() string {
	return c.name
}

// AddPoints is for implementing the Competitor Interface
func (c *SimpleCompetitor) AddPoints(p int) {
	c.GroupPoints += p
}

// GetPoints is for implementing the Competitor Interface
func (c *SimpleCompetitor) GetPoints() int {
	return c.GroupPoints
}

// ID is for implementing the Competitor Interface
func (c *SimpleCompetitor) ID() int {
	return c.id
}

// SetDrawNumber is for the Competitor Interface
func (c *SimpleCompetitor) SetDrawNumber(number int) {
	c.drawNumber = number
}
