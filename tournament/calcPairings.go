package tournament

import "strconv"

type pairing struct {
	competitor1 Competitor
	competitor2 Competitor
	round       int
}

func (p pairing) inPairings( ps []pairing) bool {
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

func calcRoundsFromPairings(pairing []pairing) int {
	result := 0
	for _, p := range pairing {
		if result < p.round {
			result = p.round
		}
	}
	return result
}

// calcPairings calcs the pairings needed when everyone needs to play against each other
func calcPairings(c []Competitor) []pairing {
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
		addToResult(mc, &result, c1, i)
		// shift
		mc = append(mc[1:], mc[0])
	}
	return result
}

func addToResult(tempCompetitors []Competitor, result *[]pairing, c1 Competitor, i int) {
	for j := 0; j < (len(tempCompetitors)/2)+1; j++ {
		if j == 0 {
			addPair(result, c1, tempCompetitors[len(tempCompetitors)-1], i+1)
		} else {
			addPair(result, tempCompetitors[j-1], tempCompetitors[len(tempCompetitors)-1-j], i+1)
		}
	}
}

func addPair(pairings *[]pairing, c1 Competitor, c2 Competitor, round int) {
	if c1.name == "" || c2.name == "" {
		return
	}
	pair := pairing{c1, c2, round}
	*pairings = append(*pairings, pair)
}
