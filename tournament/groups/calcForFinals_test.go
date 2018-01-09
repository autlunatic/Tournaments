package groups_test

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

var testCompetitors []competitors.C

func getGroupsFor1() []groups.G {
	g := make([]groups.G, 2)
	testCompetitors = competitors.NewTestCompetitors(8)
	g[0].AddCompetitors(testCompetitors[0:4])
	g[1].AddCompetitors(testCompetitors[4:8])
	return g
}
func getWantedFor1() []pairings.P {
	out := make([]pairings.P, 2)
	out[0] = pairings.P{Competitor1ID: 8, Competitor2ID: 3, Round: -2, ID: -2, GroupID: 0}
	out[1] = pairings.P{Competitor1ID: 7, Competitor2ID: 4, Round: -2, ID: -1, GroupID: 0}
	return out
}

func TestCalcPairingsForFinals(t *testing.T) {
	type args struct {
		groups        []groups.G
		finalistCount int
	}
	tests := []struct {
		name    string
		args    args
		want    []pairings.P
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "2 Groups 8 Competitors semifinals", args: args{getGroupsFor1(), 4}, want: getWantedFor1(), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := groups.CalcPairingsForFinals(tt.args.groups, tt.args.finalistCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcPairingsForFinals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcPairingsForFinals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcFinalistsPairingIDs(t *testing.T) {
	tests := []struct {
		name        string
		finalRounds int
		want        []int
	}{
		// TODO: Add test cases.
		{name: "2 finalists", finalRounds: 1, want: []int{1, 2}},
		{name: "4 finalists", finalRounds: 2, want: []int{1, 4, 3, 2}},
		{name: "8 finalists", finalRounds: 3, want: []int{1, 8, 5, 4, 3, 6, 7, 2}},
		{name: "64 finalists", finalRounds: 6, want: []int{1, 64, 33, 32, 17, 48, 49, 16, 9, 56, 41, 24, 25, 40, 57, 8, 5, 60, 37, 28, 21, 44, 53, 12, 13, 52, 45, 20, 29, 36, 61, 4, 3, 62, 35, 30, 19, 46, 51, 14, 11, 54, 43, 22, 27, 38, 59, 6, 7, 58, 39, 26, 23, 42, 55, 10, 15, 50, 47, 18, 31, 34, 63, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groups.CalcFinalistsPairingIDs(tt.finalRounds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcPairingsForFinals() = %v, want %v", got, tt.want)
			}
		})
	}

}
