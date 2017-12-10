package competitors

import "strconv"

// NewTestCompetitors generates some Competitors for testing
func NewTestCompetitors(count int) Competitors {
	competitors := Competitors{}
	for i := 1; i <= count; i++ {
		competitors.Items = append(competitors.Items, NewCompetitor(strconv.Itoa(i)))
	}
	return competitors
}
