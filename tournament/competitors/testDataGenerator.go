package competitors

import "strconv"

// NewTestCompetitors generates some Competitors for testing
func NewTestCompetitors(count int) Competitors {
	competitors := Competitors{}
	for i := 1; i <= count; i++ {
		c := NewCompetitor(strconv.Itoa(i), i-1)
		competitors.Items = append(competitors.Items, c)

	}
	return competitors
}
