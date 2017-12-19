package competitors

import (
	"reflect"
	"testing"
)

func wantedSlice() []Competitor {
	c := make([]Competitor, 4)
	c[0] = GetCompetitor(1)
	c[1] = GetCompetitor(2)
	c[2] = GetCompetitor(3)
	c[3] = GetCompetitor(0)
	return c
}

func TestGetCompetitorsSortedByGroupPoints(t *testing.T) {
	Items = NewTestCompetitors(4)
	GetCompetitor(0).AddPoints(1)
	GetCompetitor(1).AddPoints(10)
	GetCompetitor(2).AddPoints(9)
	GetCompetitor(3).AddPoints(2)

	tests := []struct {
		name string
		want []Competitor
	}{
		{"sort", wantedSlice()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompetitorsSortedByGroupPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompetitorsSortedByGroupPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
