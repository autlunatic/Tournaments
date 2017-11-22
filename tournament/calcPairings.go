package tournament

type pairing struct {
	competitor1 Competitor
	competitor2 Competitor
	round       int
}

func calcPairings(c CompetitorsGetter) []pairing {
	var result []pairing
	competitors := c.getCompetitors()
	// make a copy of the competitors
	tempCompetitors := make([]Competitor, len(competitors))
	copy(tempCompetitors, competitors)
	// if the count is odd add one dummy competitor for the roundRobin
	if len(competitors)%2 > 0 {
		tempCompetitors = append(tempCompetitors, newCompetitor(""))
	}
	// cut the first fixated competitor see roundrobin
	if len(tempCompetitors) > 1 {
		tempCompetitors = append(tempCompetitors[1:])
	}
	// shift one time so it starts with 1v2
	tempCompetitors = append(tempCompetitors[1:], tempCompetitors[0])
	c1 := competitors[0]
	for i := 0; i < len(tempCompetitors); i++ {
		AddToResult(tempCompetitors, &result, c1, i)
		// shift
		tempCompetitors = append(tempCompetitors[1:], tempCompetitors[0])
	}
	return result
}

func AddToResult(tempCompetitors []Competitor, result *[]pairing, c1 Competitor, i int) {
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
