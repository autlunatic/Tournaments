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

var tc1 []competitors.C
var tc2 []competitors.C
var tc3 []competitors.C
var tc4 []competitors.C

func generate9CompetitorsIn3Groups() detFinFields {
	tc1 = competitors.NewTestCompetitors(9)
	g := make([]G, 3)
	g[0].AddCompetitors(tc1[0:3])
	g[1].AddCompetitors(tc1[3:6])
	g[2].AddCompetitors(tc1[6:])
	return detFinFields{g, 4}
}

func generate15CompetitorsIn3Groups() detFinFields {
	tc2 = competitors.NewTestCompetitors(15)

	g := make([]G, 3)
	g[0].AddCompetitors(tc2[0:5])
	g[1].AddCompetitors(tc2[5:10])
	g[2].AddCompetitors(tc2[10:])
	return detFinFields{g, 8}
}

func generate15CompetitorsIn1Group() detFinFields {
	tc3 = competitors.NewTestCompetitors(15)
	g := make([]G, 1)
	g[0].AddCompetitors(tc3[0:])
	return detFinFields{g, 8}
}

func generate8CompetitorsIn3Groups() detFinFields {
	tc4 = competitors.NewTestCompetitors(8)
	g := make([]G, 3)
	g[0].AddCompetitors(tc4[0:3])
	g[1].AddCompetitors(tc4[3:6])
	g[2].AddCompetitors(tc4[6:])
	return detFinFields{g, 8}
}

func TestDetermineFinalists_determine(t *testing.T) {
	tests := []struct {
		name   string
		fields detFinFields
		want   []competitors.C
	}{
		{"9  competitors 3 groups 4 finalists", generate9CompetitorsIn3Groups(), []competitors.C{tc1[8], tc1[5], tc1[2], tc1[7]}},
		{"15 competitors 3 groups 8 finalists", generate15CompetitorsIn3Groups(), []competitors.C{tc2[14], tc2[9], tc2[4], tc2[13], tc2[8], tc2[3], tc2[12], tc2[7]}},
		{"15 competitors 1 groups 8 finalists", generate15CompetitorsIn1Group(), []competitors.C{tc3[14], tc3[13], tc3[12], tc3[11], tc3[10], tc3[9], tc3[8], tc3[7]}},
		{"8  competitors 3 groups 8 finalists", generate8CompetitorsIn3Groups(), []competitors.C{tc4[7], tc4[5], tc4[2], tc4[6], tc4[4], tc4[1], tc4[3], tc4[0]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetermineFinalists(tt.fields.grps, tt.fields.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetermineFinalists.determine() = %v, want %v", got, tt.want)
			}
		})
	}
}
