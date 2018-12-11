package competitors

import "strconv"

var testCompNames = [...]string{
	"Benni",
	"Dani",
	"Mona",
	"Andrea",
	"Zoé",
	"Andreas",
	"Bernhard",
	"Florian",
	"Simon",
}

// NewTestCompetitors generates some Competitors for testing
func NewTestCompetitors(count int) []C {
	competitors := make([]C, count)
	for i := 1; i <= count; i++ {
		var compName string

		if len(testCompNames) >= i {
			compName = testCompNames[i-1]
		} else {
			compName = strconv.Itoa(i)
		}
		c := New(compName, i-1)
		competitors[i-1] = c
		c.AddResult(ResultPoints{i, 0, i, 0, -1})
		c.SetDrawNumber(i)
	}
	return competitors
}
