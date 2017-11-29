package tournament

import "strconv"

type pairing struct {
	competitor1 Competitor
	competitor2 Competitor
	round       int
	id          int
	groupId     int
}

type pairings []pairing

func (p pairing) inPairings(ps []pairing) bool {
	for _, mp := range ps {
		if mp.Equals(p) {
			return true
		}
	}
	return false
}

func (p pairing) Equals(p2 pairing) bool {
	return p.competitor1 == p2.competitor1 &&
		p.competitor2 == p2.competitor2 &&
		p.round == p2.round
}

func (p pairing) toString() string {
	return "round: " + strconv.Itoa(p.round) + "; " + p.competitor1.name + " vs. " + p.competitor2.name
}

func getMaxRoundOfPairings(pairing []pairing) int {
	result := 0
	for _, p := range pairing {
		if result < p.round {
			result = p.round
		}
	}
	return result
}

type sortByRound pairings

func (a sortByRound) Len() int      { return len(a) }
func (a sortByRound) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortByRound) Less(i, j int) bool {
	if a[i].round != a[j].round {
		return a[i].round < a[j].round
	}
	if  a[i].groupId != a[j].groupId{
		return a[i].groupId < a[j].groupId
	}
	return a[i].competitor1.drawNumber < a[j].competitor1.drawNumber
}

// calcPairings calculates the pairings needed when everyone needs to play against each other
func calcPairings(c []Competitor, groupId int) []pairing {
	var result []pairing
	competitors := c
	// make a copy of the competitors
	mc := make([]Competitor, len(competitors))
	copy(mc, competitors)
	// if the count is odd add one dummy competitor for the roundRobin
	if len(competitors)%2 > 0 {
		mc = append(mc, newCompetitor(""))
	}
	// cut the first fixated competitor see roundrobin
	if len(mc) > 1 {
		mc = append(mc[1:])
	}
	// shift one time so it starts with 1v2
	mc = append(mc[1:], mc[0])
	c1 := competitors[0]
	for i := 0; i < len(mc); i++ {
		addToResult(mc, &result, c1, i, groupId)
		// shift
		mc = append(mc[1:], mc[0])
	}
	return result
}

func addToResult(c []Competitor, result *[]pairing, c1 Competitor, i int, groupid int) {
	for j := 0; j < (len(c)/2)+1; j++ {
		if j == 0 {
			addPair(result, c1, c[len(c)-1], i+1, groupid)
		} else {
			addPair(result, c[j-1], c[len(c)-1-j], i+1, groupid)
		}
	}
}

func addPair(pairings *[]pairing, c1 Competitor, c2 Competitor, round int ,groupId int) {
	if c1.name == "" || c2.name == "" {
		return
	}
	pair := pairing{c1, c2, round, 0, groupId}
	*pairings = append(*pairings, pair)
}
