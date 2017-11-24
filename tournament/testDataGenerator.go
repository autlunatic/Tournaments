package tournament

import "strconv"

func newTestCompetitors(count int) Competitors {
	competitors := Competitors{}
	for i := 1; i <= count; i++ {
		competitors.items = append(competitors.items, newCompetitor(strconv.Itoa(i)))
	}
	return competitors
}
