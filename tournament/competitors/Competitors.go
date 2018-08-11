package competitors

import (
	"errors"
	"strings"
)

// GetCompetitor gets the Competitor of the Items with the given ID
func GetCompetitor(c []C, ID int) C {
	for _, item := range c {
		if item.ID() == ID {
			return item
		}
	}
	return &EmptyCompetitor{}
}

// New generates a New SimpleCompetitor with the given Name
func New(name string, id int) *SimpleCompetitor {
	c := new(SimpleCompetitor)
	c.TeamName = name
	c.Id = id
	return c
}

// ClearPoints sets the Points of all Items to zero
func ClearPoints(c []C) {
	for _, c := range c {
		c.ClearResults()
	}
}

// ResultPoints represents one Result for the competitor
type ResultPoints struct {
	GamePoints          int
	GamePointsNegative  int
	GroupPoints         int
	GroupPointsNegative int
	AgainstCompetitorID int
}

// C is the interface used by the Tournament package for Operations with the Competitor
type C interface {
	AddResult(result ResultPoints)
	GetResults() []ResultPoints
	ClearResults()
	ID() int
	Name() string
	DrawNumber() int
	SetDrawNumber(int)
	SetGroupPlacement(int)
	GroupPlacement() int
}

// SumResultPoints calcs the sum of all points and adds the component ids of the competitor which are defeated
func SumResultPoints(rps []ResultPoints) (outPoints ResultPoints, wonAgainst []int) {

	for _, rp := range rps {
		outPoints.GamePoints += rp.GamePoints
		outPoints.GamePointsNegative += rp.GamePointsNegative
		outPoints.GroupPoints += rp.GroupPoints
		outPoints.GroupPointsNegative += rp.GroupPointsNegative

		if rp.GroupPoints > rp.GroupPointsNegative {
			wonAgainst = append(wonAgainst, rp.AgainstCompetitorID)
		}
	}
	return outPoints, wonAgainst
}

// SimpleCompetitor holds information for an minimalistic Competitor
type SimpleCompetitor struct {
	Id          int
	TeamName    string
	GroupPoints int
	GamePoints  int
	DrawNr      int
	GroupPlace  int
	Results     []ResultPoints
}

// ContainsName Checks if the competitor Name is already taken
func ContainsName(cs []C, name string) bool {
	for _, c := range cs {
		if strings.ToUpper(c.Name()) == strings.ToUpper(name) {
			return true
		}
	}
	return false
}

// NameToID returns the competitor ID of the given Competitor Name
func NameToID(cs []C, name string) int {
	for _, c := range cs {
		if c.Name() == name {
			return c.ID()
		}
	}
	return -1
}

// GetMaxID returns the highest ID of a competitor in the given slice
func GetMaxID(cs []C) int {
	var out int
	for _, c := range cs {
		if out < c.ID() {
			out = c.ID()
		}
	}
	return out
}

// AddByName adds a competitor by name, the ID is automatically generated
// it returns an error if the name is already taken
func AddByName(cs []C, name string) ([]C, error) {
	if ContainsName(cs, name) {
		return cs, errors.New("Name already taken")
	}
	cs = append(cs, New(name, GetMaxID(cs)+1))
	return cs, nil
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
func (c *SimpleCompetitor) DrawNumber() int {
	return c.DrawNr
}

// AddResult is for implementing the Competitor Interface
func (c *SimpleCompetitor) AddResult(result ResultPoints) {
	c.Results = append(c.Results, result)
}

// GetResults is for implementing the competitor interface
func (c *SimpleCompetitor) GetResults() []ResultPoints {
	return c.Results
}

// ClearResults is for the Competitor Interface
func (c *SimpleCompetitor) ClearResults() {
	c.Results = []ResultPoints{}
}

// Name is for implementing the Competitor Interface
func (c *SimpleCompetitor) Name() string {
	return c.TeamName
}

// ID is for implementing the Competitor Interface
func (c *SimpleCompetitor) ID() int {
	return c.Id
}

// SetDrawNumber is for the Competitor Interface
func (c *SimpleCompetitor) SetDrawNumber(number int) {
	c.DrawNr = number
}

// SetGroupPlacement is for the Competitor Interface
func (c *SimpleCompetitor) SetGroupPlacement(in int) {
	c.GroupPlace = in
}

// GroupPlacement is for the Competitor Interface
func (c *SimpleCompetitor) GroupPlacement() int {
	return c.GroupPlace
}

// EmptyCompetitor is used as Return when the competitor can not be found
type EmptyCompetitor struct{}

// DrawNumber is for implementing the Competitor interface
func (c *EmptyCompetitor) DrawNumber() int {
	return 0
}

// Name is for implementing the Competitor Interface
func (c *EmptyCompetitor) Name() string {
	return ""
}

// AddResult is for implementing the Competitor Interface
func (c *EmptyCompetitor) AddResult(result ResultPoints) {
}

// GetResults is for implementing the competitor interface
func (c *EmptyCompetitor) GetResults() []ResultPoints {
	return []ResultPoints{}
}

// ClearResults is for the Competitor Interface
func (c *EmptyCompetitor) ClearResults() {

}

// ID is for implementing the Competitor Interface
func (c *EmptyCompetitor) ID() int {
	return -1
}

// SetDrawNumber is for the Competitor Interface
func (c *EmptyCompetitor) SetDrawNumber(number int) {
}

// SetGroupPlacement is for the Competitor Interface
func (c *EmptyCompetitor) SetGroupPlacement(in int) {
}

// GroupPlacement is for the Competitor Interface
func (c *EmptyCompetitor) GroupPlacement() int {
	return 0
}
