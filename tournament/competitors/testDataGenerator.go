package competitors

import "strconv"

// NewTestCompetitors generates some Competitors for testing
func NewTestCompetitors(count int) []C {
	competitors := make([]C, count)
	for i := 1; i <= count; i++ {
		c := New(strconv.Itoa(i), i-1)
		competitors[i-1] = c
		c.AddPoints(i)
	}
	return competitors
}
