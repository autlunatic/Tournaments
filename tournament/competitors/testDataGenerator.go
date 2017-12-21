package competitors

import "strconv"

// NewTestCompetitors generates some Competitors for testing
func NewTestCompetitors(count int) []Competitor {
	competitors := make([]Competitor, count)
	for i := 1; i <= count; i++ {
		c := NewCompetitor(strconv.Itoa(i), i-1)
		competitors[i-1] = c
	}
	return competitors
}
