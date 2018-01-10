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
	g := groups.NewTestGroups(2)
	testCompetitors = competitors.NewTestCompetitors(8)
	g[0].AddCompetitors(testCompetitors[0:4])
	g[1].AddCompetitors(testCompetitors[4:8])
	return g
}

func getGroupsFor2(addPointsTo1 bool) []groups.G {
	g := groups.NewTestGroups(3)
	testCompetitors = competitors.NewTestCompetitors(12)
	if addPointsTo1 {
		testCompetitors[4].AddPoints(15)
	}
	g[0].AddCompetitors(testCompetitors[0:4])
	g[1].AddCompetitors(testCompetitors[4:8])
	g[2].AddCompetitors(testCompetitors[9:])
	return g
}

func getWantedFor1() []pairings.P {
	out := make([]pairings.P, 2)
	out[0] = pairings.P{Competitor1ID: 7, Competitor2ID: 2, Round: -4, ID: -1, GroupID: 0}
	out[1] = pairings.P{Competitor1ID: 6, Competitor2ID: 3, Round: -4, ID: -2, GroupID: 0}
	return out
}

func getWantedFor2() []pairings.P {
	// 11, 7, 3,10, 6, 2, 9, 5
	//  1, 8, 5, 4, 3, 6, 7, 2
	out := make([]pairings.P, 4)
	out[0] = pairings.P{Competitor1ID: 11, Competitor2ID: 5, Round: -8, ID: -1, GroupID: 0}
	out[1] = pairings.P{Competitor1ID: 2, Competitor2ID: 10, Round: -8, ID: -2, GroupID: 0}
	out[2] = pairings.P{Competitor1ID: 3, Competitor2ID: 6, Round: -8, ID: -3, GroupID: 0}
	out[3] = pairings.P{Competitor1ID: 9, Competitor2ID: 7, Round: -8, ID: -4, GroupID: 0}
	return out
}

func getWantedFor3() []pairings.P {
	//4,11, 3, 10, 7, 2, 9, 6
	//1, 8, 5,  4, 3, 6, 7, 2
	out := make([]pairings.P, 4)
	out[0] = pairings.P{Competitor1ID: 4, Competitor2ID: 9, Round: -8, ID: -1, GroupID: 0}
	out[1] = pairings.P{Competitor1ID: 2, Competitor2ID: 10, Round: -8, ID: -2, GroupID: 0}
	out[2] = pairings.P{Competitor1ID: 3, Competitor2ID: 7, Round: -8, ID: -3, GroupID: 0}
	out[3] = pairings.P{Competitor1ID: 6, Competitor2ID: 11, Round: -8, ID: -4, GroupID: 0}
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
		{name: "2 Groups 4 Competitors semifinals", args: args{getGroupsFor1(), 4}, want: getWantedFor1(), wantErr: false},
		{name: "3 Groups 8 Competitors quarterfinals", args: args{getGroupsFor2(false), 8}, want: getWantedFor2(), wantErr: false},
		{name: "3 Groups 8 Competitors quarterfinals not first round not from group", args: args{getGroupsFor2(true), 8}, want: getWantedFor3(), wantErr: false},
		//TODO Need an additional test when only last one fails because of same grp
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
		{name: "2 finalists", finalRounds: 1, want: []int{1, 2}},
		{name: "4 finalists", finalRounds: 2, want: []int{1, 4, 3, 2}},
		{name: "8 finalists", finalRounds: 3, want: []int{1, 8, 5, 4, 3, 6, 7, 2}},
		{name: "64 finalists", finalRounds: 6, want: []int{1, 64, 33, 32, 17, 48, 49, 16, 9, 56, 41, 24, 25, 40, 57, 8, 5, 60, 37, 28, 21, 44, 53, 12, 13, 52, 45, 20, 29, 36, 61, 4, 3, 62, 35, 30, 19, 46, 51, 14, 11, 54, 43, 22, 27, 38, 59, 6, 7, 58, 39, 26, 23, 42, 55, 10, 15, 50, 47, 18, 31, 34, 63, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groups.CalcFinalistRankings(tt.finalRounds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcPairingsForFinals() = %v, want %v", got, tt.want)
			}
		})
	}

}
