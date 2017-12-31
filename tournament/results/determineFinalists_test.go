package results

import (
	"fmt"
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

func generate15CompetitorsIn3Groups() fields {
	testCompetitors = competitors.NewTestCompetitors(15)
	for i := 0; i < 15; i++ {
		testCompetitors[i].AddPoints(i)
	}
	g := make([]groups.Group, 3)
	g[0].AddCompetitors(testCompetitors[0:5])
	g[1].AddCompetitors(testCompetitors[5:10])
	g[2].AddCompetitors(testCompetitors[10:])
	return fields{testCompetitors, g, 8}
}

func generate9CompetitorsIn3Groups() fields {
	testCompetitors = competitors.NewTestCompetitors(9)
	for i := 0; i < 9; i++ {
		testCompetitors[i].AddPoints(i)
	}
	g := make([]groups.Group, 3)
	g[0].AddCompetitors(testCompetitors[0:3])
	g[1].AddCompetitors(testCompetitors[3:6])
	g[2].AddCompetitors(testCompetitors[6:])
	return fields{testCompetitors, g, 4}
}
func generateWantedCompetitorsFirstTest() []competitors.Competitor {
	out := make([]competitors.Competitor, 4)
	out[0] = testCompetitors[8]
	out[1] = testCompetitors[5]
	out[2] = testCompetitors[2]
	out[3] = testCompetitors[7]
	return out
}
func generateWantedCompetitorsSecondTest() []competitors.Competitor {
	out := make([]competitors.Competitor, 8)
	out[0] = testCompetitors[14]
	out[1] = testCompetitors[9]
	out[2] = testCompetitors[4]
	out[3] = testCompetitors[13]
	out[4] = testCompetitors[8]
	out[5] = testCompetitors[3]
	out[6] = testCompetitors[12]
	out[7] = testCompetitors[7]
	return out
}
func TestDetermineFinalists_determine(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   []competitors.Competitor
	}{
		{"9 competitors 3 groups 4 finalists", generate9CompetitorsIn3Groups(), generateWantedCompetitorsFirstTest()},
		{"15 competitors 3 groups 8 finalists", generate15CompetitorsIn3Groups(), generateWantedCompetitorsSecondTest()},
	}
	for _, tt := range tests {
		fmt.Println("start test: ", tt.name)
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
