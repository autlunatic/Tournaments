package groups

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

var testCompetitors []competitors.C

func getGroupsFor1() []G {
	g := make([]G, 2)
	testCompetitors = competitors.NewTestCompetitors(8)
	g[0].AddCompetitors(testCompetitors[0:4])
	g[1].AddCompetitors(testCompetitors[4:8])
	return g
}
func getWantedFor1() []pairings.P {
	out := make([]pairings.P, 2)
	out[0] = pairings.P{Competitor1ID: 4, Competitor2ID: 7, Round: -2, ID: -1, GroupID: 0}
	out[1] = pairings.P{Competitor1ID: 3, Competitor2ID: 8, Round: -2, ID: -2, GroupID: 0}
	return out
}

func TestCalcForFinals(t *testing.T) {
	type args struct {
		groups []G
	}
	tests := []struct {
		name string
		args args
		want []pairings.P
	}{
		{name: "2 Groups 8 Competitors semifinals", args: args{getGroupsFor1()}, want: getWantedFor1()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcForFinals(tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcForFinals() = %v, want %v", got, tt.want)
			}
		})
	}
}
