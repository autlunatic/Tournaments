package competitors

import (
	"errors"
)

// GetCompetitor gets the Competitor of the Items with the given ID
func GetCompetitor(c []C, ID int) C {
	for _, item := range c {
		if item.ID() == ID {
			return item
		}
	}
	return nil
}

// New generates a New SimpleCompetitor with the given Name
func New(name string, id int) *SimpleCompetitor {
	c := new(SimpleCompetitor)
	c.name = name
	c.id = id
	return c
}

// ClearPoints sets the Points of all Items to zero
func ClearPoints(c []C) {
	for _, c := range c {
		c.AddPoints(-c.GetPoints())
	}
}

// C is the interface used by the Tournament package for Operations with the Competitor
type C interface {
	AddPoints(p int)
	GetPoints() int
	ID() int
	Name() string
	DrawNumber() int64
	SetDrawNumber(int64)
	SetGroupPlacement(int)
	GroupPlacement() int
}

// SimpleCompetitor holds information for an minimalistic Competitor
type SimpleCompetitor struct {
	id             int
	name           string
	GroupPoints    int
	drawNumber     int64
	groupPlacement int
}

// Add adds a competitor to the slice of competitors
func Add(cs []C, c C) ([]C, error) {
	for _, item := range cs {
		if item.ID() == c.ID() {
			return cs, errors.New("ID is already taken")
		}
	}
	cs = append(cs, c)
	return cs, nil
}

// Delete deletes the Competitor with the given ID
// it returns a count of Deleted rows, which in a normal use case should always be 0 or one
func Delete(cs []C, competitorID int) int {
	var deleteCount int
	for i := len(cs) - 1; i >= 0; i-- {
		if cs[i].ID() == competitorID {
			cs = append(cs[:i], cs[i+1:]...)
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

// SetGroupPlacement is for the Competitor Interface
func (c *SimpleCompetitor) SetGroupPlacement(in int) {
	c.groupPlacement = in
}

// GroupPlacement is for the Competitor Interface
func (c *SimpleCompetitor) GroupPlacement() int {
	return c.groupPlacement
}
