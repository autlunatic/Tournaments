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
