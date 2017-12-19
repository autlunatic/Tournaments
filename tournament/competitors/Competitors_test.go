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

func TestDelete(t *testing.T) {
	Items = NewTestCompetitors(5)
	type args struct {
		competitorID int
	}
	tests := []struct {
		name            string
		arg             args
		want            int
		competitorToAdd Competitor
	}{
		{name: "Competitor not in items", arg: args{competitorID: 9}, want: 0, competitorToAdd: nil},
		{name: "Competitor one in items", arg: args{competitorID: 3}, want: 1, competitorToAdd: nil},
		{"Competitor two in items", args{competitorID: 1}, 2, NewCompetitor("1", 1)},
	}
	for _, tt := range tests {
		if tt.competitorToAdd != nil {
			Items.Items = append(Items.Items, tt.competitorToAdd)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.arg.competitorID); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
