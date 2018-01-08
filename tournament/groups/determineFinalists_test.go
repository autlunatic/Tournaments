package groups

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

type detFinFields struct {
	grps  []G
	count int
}

var tc []competitors.C

func generate15CompetitorsIn3Groups() detFinFields {
	tc = competitors.NewTestCompetitors(15)
	for i := 0; i < 15; i++ {
		tc[i].AddPoints(i)
	}
	g := make([]G, 3)
	g[0].AddCompetitors(tc[0:5])
	g[1].AddCompetitors(tc[5:10])
	g[2].AddCompetitors(tc[10:])
	return detFinFields{g, 8}
}

func generate9CompetitorsIn3Groups() detFinFields {
	tc = competitors.NewTestCompetitors(9)
	for i := 0; i < 9; i++ {
		tc[i].AddPoints(i)
	}
	g := make([]G, 3)
	g[0].AddCompetitors(tc[0:3])
	g[1].AddCompetitors(tc[3:6])
	g[2].AddCompetitors(tc[6:])
	return detFinFields{g, 4}
}
func generateWantedCompetitorsFirstTest() []competitors.C {
	out := make([]competitors.C, 4)
	out[0] = tc[8]
	out[1] = tc[5]
	out[2] = tc[2]
	out[3] = tc[7]
	return out
}
func generateWantedCompetitorsSecondTest() []competitors.C {
	out := make([]competitors.C, 8)
	out[0] = tc[14]
	out[1] = tc[9]
	out[2] = tc[4]
	out[3] = tc[13]
	out[4] = tc[8]
	out[5] = tc[3]
	out[6] = tc[12]
	out[7] = tc[7]
	return out
}
func TestDetermineFinalists_determine(t *testing.T) {
	tests := []struct {
		name   string
		fields detFinFields
		want   []competitors.C
	}{
		{"9 competitors 3 groups 4 finalists", generate9CompetitorsIn3Groups(), generateWantedCompetitorsFirstTest()},
		{"15 competitors 3 groups 8 finalists", generate15CompetitorsIn3Groups(), generateWantedCompetitorsSecondTest()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetermineFinalists(tt.fields.grps, tt.fields.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetermineFinalists.determine() = %v, want %v", got, tt.want)
			}
		})
	}
}
