package competitors

import (
	"errors"
	"sort"
)

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
	for _, item := range Items.Items {
		if item.ID() == ID {
			return item
		}
	}
	return nil
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

// GetCompetitorsSortedByGroupPoints returns a slice of Competitor which is sorted by GroupPoints ;)
func GetCompetitorsSortedByGroupPoints() []Competitor {
	sorter := &sortByGroupPoints{Items.Items}
	sort.Sort(sorter)
	return sorter.items
}

// ClearPoints sets the Points of all Items to zero
func ClearPoints() {
	for _, c := range Items.Items {
		c.AddPoints(-c.GetPoints())
	}
}

// Competitor is the interface used by the Tournament package for Operations with the Competitor
type Competitor interface {
	AddPoints(p int)
	GetPoints() int
	ID() int
	Name() string
	DrawNumber() int64
	SetDrawNumber(int64)
}

// SimpleCompetitor holds information for an minimalistic Competitor
type SimpleCompetitor struct {
	id          int
	name        string
	GroupPoints int
	drawNumber  int64
}

// Add adds a competitor to the slice of competitors
func Add(c Competitor) error {
	for _, item := range Items.Items {
		if item.ID() == c.ID() {
			return errors.New("ID is already taken")
		}
	}
	Items.Items = append(Items.Items, c)
	return nil
}

// Delete deletes the Competitor with the given ID
// it returns a count of Deleted rows, which in a normal use case should always be 0 or one
func Delete(competitorID int) int {
	var deleteCount int
	for i := len(Items.Items) - 1; i >= 0; i-- {
		if Items.Items[i].ID() == competitorID {
			Items.Items = append(Items.Items[:i], Items.Items[i+1:]...)
			deleteCount++
		}
	}
	return deleteCount
}

// DrawNumber is for implementing the Competitor interface
func (c *SimpleCompetitor) DrawNumber() int64 {
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
func (c *SimpleCompetitor) SetDrawNumber(number int64) {
	c.drawNumber = number
}
