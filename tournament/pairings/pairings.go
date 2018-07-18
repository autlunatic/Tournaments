package pairings

import (
	"errors"
	"sort"
	"strconv"
	"time"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// P holds information about one game,
type P struct {
	Competitor1ID int
	Competitor2ID int
	Round         int
	ID            int
	GroupID       int
	StartTime     time.Time
	Court         int
}

// InPairings checks if the given slice of Pairing contains the Pairing
func (p P) InPairings(ps []P) bool {
	for _, mp := range ps {
		if mp.equals(p) {
			return true
		}
	}
	return false
}

// OfCompetitorID returns all Pairings where the competitor is playing
func OfCompetitorID(ps []P, compID int) []P {
	var out []P
	for i, p := range ps {
		if p.Competitor1ID == compID || p.Competitor2ID == compID {
			out = append(out, ps[i])
		}
	}
	return out
}

func (p P) equals(p2 P) bool {
	return p.Competitor1ID == p2.Competitor1ID &&
		p.Competitor2ID == p2.Competitor2ID &&
		p.Round == p2.Round
}

// ToString provides a simple string representation of the pairing
func (p P) ToString() string {
	return "round: " + strconv.Itoa(p.Round) + "; " + strconv.Itoa(p.Competitor1ID) + " vs. " + strconv.Itoa(p.Competitor2ID)
}

// GetMaxRoundOfPairings returns the max Round - number in the slice of Pairing
func GetMaxRoundOfPairings(pairing []P) int {
	result := 0
	for _, p := range pairing {
		if result < p.Round {
			result = p.Round
		}
	}
	return result
}

// SortByRoundGroupDrawnumber implements the Interface interface for sorting
type SortByRoundGroupDrawnumber struct {
	Pairs []P
	Comps []competitors.C
}

func (a SortByRoundGroupDrawnumber) Len() int      { return len(a.Pairs) }
func (a SortByRoundGroupDrawnumber) Swap(i, j int) { a.Pairs[i], a.Pairs[j] = a.Pairs[j], a.Pairs[i] }
func (a SortByRoundGroupDrawnumber) Less(i, j int) bool {
	if a.Pairs[i].Round != a.Pairs[j].Round {
		return a.Pairs[i].Round < a.Pairs[j].Round
	}
	if a.Pairs[i].GroupID != a.Pairs[j].GroupID {
		return a.Pairs[i].GroupID < a.Pairs[j].GroupID
	}
	return competitors.GetCompetitor(a.Comps, a.Pairs[i].Competitor1ID).DrawNumber() <
		competitors.GetCompetitor(a.Comps, a.Pairs[j].Competitor1ID).DrawNumber()

}

// CalcPairings calculates the pairings needed when everyone needs to play against each other
func CalcPairings(c []competitors.C, groupID int) ([]P, error) {
	var result []P

	if len(c) == 0 {
		return result, errors.New("no Competitors given, cannot calc pairings")
	}
	cs := c
	// make a copy of the cs
	mc := make([]competitors.C, len(cs))
	copy(mc, cs)
	// if the count is odd add one dummy competitor for the roundRobin
	if len(cs)%2 > 0 {
		mc = append(mc, competitors.New("", -1))
	}
	// cut the first fixated competitor see roundrobin
	if len(mc) > 1 {
		mc = append(mc[1:])
	}
	// shift one time so it starts with 1v2
	mc = append(mc[1:], mc[0])
	c1 := cs[0]
	for i := 0; i < len(mc); i++ {
		addToResult(mc, &result, c1.ID(), i, groupID)
		// shift
		mc = append(mc[1:], mc[0])
	}
	var s SortByRoundGroupDrawnumber
	s.Comps = c
	s.Pairs = result
	sort.Sort(s)
	return s.Pairs, nil
}

func addToResult(c []competitors.C, result *[]P, c1 int, i int, groupID int) {
	for j := 0; j < (len(c)/2)+1; j++ {
		if j == 0 {
			addPair(result, c1, c[len(c)-1].ID(), i+1, groupID)
		} else {
			addPair(result, c[j-1].ID(), c[len(c)-1-j].ID(), i+1, groupID)
		}
	}
}

func addPair(pairings *[]P, c1 int, c2 int, round int, groupID int) {
	if c1 == -1 || c2 == -1 {
		return
	}
	pair := P{c1, c2, round, 0, groupID, time.Time{}, -1}

	*pairings = append(*pairings, pair)
}

// LatestGameStart returns the starttime of the LastGame played
func LatestGameStart(ps []P) time.Time {
	var out time.Time
	for _, p := range ps {
		if p.StartTime.After(out) {
			out = p.StartTime
		}
	}
	return out
}
