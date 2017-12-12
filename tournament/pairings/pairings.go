package pairings

import (
	"errors"
	"strconv"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// Pairing holds information about one game,
type Pairing struct {
	Competitor1 competitors.Competitor
	Competitor2 competitors.Competitor
	Round       int
	ID          int
	GroupID     int
}

func (p *Pairing) addGroupPoints(cp1 int, cp2 int) {
	p.Competitor1.GroupPoints += cp1
	p.Competitor2.GroupPoints += cp2
}

// InPairings checks if the given slice of Pairing contains the Pairing
func (p Pairing) InPairings(ps []Pairing) bool {
	for _, mp := range ps {
		if mp.equals(p) {
			return true
		}
	}
	return false
}

func (p Pairing) equals(p2 Pairing) bool {
	return p.Competitor1 == p2.Competitor1 &&
		p.Competitor2 == p2.Competitor2 &&
		p.Round == p2.Round
}

// toString provides a simple string representation of the pairing
func (p Pairing) toString() string {
	return "round: " + strconv.Itoa(p.Round) + "; " + p.Competitor1.Name + " vs. " + p.Competitor2.Name
}

// GetMaxRoundOfPairings returns the max Round - number in the slice of Pairing
func GetMaxRoundOfPairings(pairing []Pairing) int {
	result := 0
	for _, p := range pairing {
		if result < p.Round {
			result = p.Round
		}
	}
	return result
}

// SortByRound implements the Interface interface for sorting
type SortByRound []Pairing

func (a SortByRound) Len() int      { return len(a) }
func (a SortByRound) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByRound) Less(i, j int) bool {
	if a[i].Round != a[j].Round {
		return a[i].Round < a[j].Round
	}
	if a[i].GroupID != a[j].GroupID {
		return a[i].GroupID < a[j].GroupID
	}
	return a[i].Competitor1.DrawNumber < a[j].Competitor1.DrawNumber
}

// CalcPairings calculates the pairings needed when everyone needs to play against each other
func CalcPairings(c []competitors.Competitor, groupID int) ([]Pairing, error) {
	var result []Pairing

	if len(c) == 0 {
		return result, errors.New("no Competitors given, cannot calc pairings!")
	}
	cs := c
	// make a copy of the cs
	mc := make([]competitors.Competitor, len(cs))
	copy(mc, cs)
	// if the count is odd add one dummy competitor for the roundRobin
	if len(cs)%2 > 0 {
		mc = append(mc, competitors.NewCompetitor(""))
	}
	// cut the first fixated competitor see roundrobin
	if len(mc) > 1 {
		mc = append(mc[1:])
	}
	// shift one time so it starts with 1v2
	mc = append(mc[1:], mc[0])
	c1 := cs[0]
	for i := 0; i < len(mc); i++ {
		addToResult(mc, &result, c1, i, groupID)
		// shift
		mc = append(mc[1:], mc[0])
	}
	return result, nil
}

func addToResult(c []competitors.Competitor, result *[]Pairing, c1 competitors.Competitor, i int, groupID int) {
	for j := 0; j < (len(c)/2)+1; j++ {
		if j == 0 {
			addPair(result, c1, c[len(c)-1], i+1, groupID)
		} else {
			addPair(result, c[j-1], c[len(c)-1-j], i+1, groupID)
		}
	}
}

func addPair(pairings *[]Pairing, c1 competitors.Competitor, c2 competitors.Competitor, round int, groupID int) {
	if c1.Name == "" || c2.Name == "" {
		return
	}
	pair := Pairing{c1, c2, round, 0, groupID}
	*pairings = append(*pairings, pair)
}
