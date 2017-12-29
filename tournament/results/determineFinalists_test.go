package results

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
)

func generate9CompetitorsIn3GroupsHalfFinals() ([]competitors.Competitor, []groups.Group) {
	c := competitors.NewTestCompetitors(9)
	g := make([]groups.Group, 3, 3)
	for i := 0; i < 9; i++ {
		c[i].AddPoints(i)
	}
	return c, g
}
func Test_determineFinalists(t *testing.T) {
	type args struct {
		in0 []competitors.Competitor
		in1 []groups.Group
	}
	tests := []struct {
		name string
		args args
		want []competitors.Competitor
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineFinalists(tt.args.in0, tt.args.in1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determineFinalists() = %v, want %v", got, tt.want)
			}
		})
	}
}
