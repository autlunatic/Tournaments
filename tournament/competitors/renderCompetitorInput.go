package competitors

import (
	"strconv"
)

// CompetitorInfo is used for displaying the competitor Name in API oder HTML
type CompetitorInfo struct {
	ID         int
	Name       string
	DrawNumber string
}

type inputCompetitors struct {
	C       []C
	ErrHTML string
}

// ToCompetitorInfo converts the competitors for output or API
func ToCompetitorInfo(c []C) []CompetitorInfo {
	var out []CompetitorInfo
	for _, ci := range c {
		out = append(out, CompetitorInfo{ID: ci.ID(),
			Name:       ci.Name(),
			DrawNumber: strconv.Itoa(ci.DrawNumber())})
	}
	return out
}
