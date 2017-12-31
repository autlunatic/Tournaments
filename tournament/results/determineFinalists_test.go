package results

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
)

type fields struct {
	comps []competitors.Competitor
	grps  []groups.Group
	count int
}

var testCompetitors []competitors.Competitor

func generate9CompetitorsIn3Groups(finalistCount int) (f fields) {
	testCompetitors = competitors.NewTestCompetitors(9)
	g := make([]groups.Group, 3)
	for i := 0; i < 9; i++ {
		testCompetitors[i].AddPoints(i)
		if finalistCount == 8 && i == 0 {

			testCompetitors[i].AddPoints(9)
		}
	}
	g[0].AddCompetitors(testCompetitors[0:3])
	g[1].AddCompetitors(testCompetitors[3:6])
	g[2].AddCompetitors(testCompetitors[6:])
	comps := make([]competitors.Competitor, len(testCompetitors))
	copy(comps, testCompetitors)
	f.comps = comps
	f.grps = g
	f.count = finalistCount
	return
}
func generateWantedCompetitorsFirstTest() []competitors.Competitor {
	out := make([]competitors.Competitor, 4)
	out[0] = testCompetitors[8]
	out[1] = testCompetitors[7]
	out[2] = testCompetitors[5]
	out[3] = testCompetitors[2]
	return out
}
func generateWantedCompetitorsSecondTest() []competitors.Competitor {
	out := make([]competitors.Competitor, 8)
	var outIndex int
	for t, c := range testCompetitors {
		if c.Name() != "2" {
			out[outIndex] = testCompetitors[t]
			outIndex++
		}
	}
	out = competitors.GetCompetitorsSortedByGroupPoints(out)
	return out
}
func TestDetermineFinalists_determine(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   []competitors.Competitor
	}{
		{"9 competitors 3 groups 4 finalists", generate9CompetitorsIn3Groups(4), generateWantedCompetitorsFirstTest()},
		{"9 competitors 3 groups 8 finalists", generate9CompetitorsIn3Groups(8), generateWantedCompetitorsSecondTest()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DetermineFinalists{
				comps: tt.fields.comps,
				grps:  tt.fields.grps,
				count: tt.fields.count,
			}
			if got := d.determine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetermineFinalists.determine() = %v, want %v", got, tt.want)
			}
		})
	}
}
