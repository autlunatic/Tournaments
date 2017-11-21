package tournament

type pairing struct {
	competitor1 Competitor
	competitor2 Competitor
	round int
}

type getCompetitors interface{
	getCompetitors() []Competitor
}

func calcPairings(c getCompetitors) []pairing {
	var result []pairing
	competitors := c.getCompetitors()
	// make a copy of the competitors
	tempCompetitors := make([]Competitor, len(competitors))
	copy(tempCompetitors, competitors)
	// if the count is odd add one dummy competitor for the roundRobin
	if len(competitors)%2 > 0 {
		tempCompetitors = append(tempCompetitors, Competitor{""})
	}
	// cut the first fixated competitor see roundrobin
	if len(tempCompetitors) > 1 {
		tempCompetitors = append(tempCompetitors[1:])
	}
	// shift one time so it starts with 1v2
	tempCompetitors = append(tempCompetitors[1:], tempCompetitors[0])
	for i := 0; i < len(tempCompetitors); i++ {
		c1 := competitors[0]

		for j := 0; j < (len(tempCompetitors)/2)+1; j++ {
			if j == 0 {
				addPair(&result, c1, tempCompetitors[len(tempCompetitors)-1], i+1)
			} else {
				addPair(&result, tempCompetitors[j-1], tempCompetitors[len(tempCompetitors)-1-j], i+1)
			}
		}
		// shift
		tempCompetitors = append(tempCompetitors[1:], tempCompetitors[0])
	}
	return result
}

func addPair(pairings *[]pairing, c1 Competitor, c2 Competitor, round  int) {
	if c1.name == "" || c2.name == "" {
		return
	}
	pair := pairing{c1, c2, round}
	*pairings = append(*pairings, pair)
}
