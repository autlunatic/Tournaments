package groups

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func calcGroupsForTest() []G {
	c := competitors.NewTestCompetitors(9)
	g, _ := CalcGroups(c, 3)
	for gi := range g {
		for _, i := range g[gi].Competitors {

			fmt.Println(i)
		}
	}
	return g
}

func getGroupInfoWanted() []GroupInfo {
	gi := make([]GroupInfo, 3)
	gi[0] = GroupInfo{1, []string{"Benni", "Dani", "Mona"}}
	gi[1] = GroupInfo{2, []string{"Andrea", "Zoé", "Andreas"}}
	gi[2] = GroupInfo{3, []string{"Bernhard", "Florian", "Simon"}}
	return gi
}
func TestGToGroupInfo(t *testing.T) {
	type args struct {
		c []competitors.C
		g []G
	}
	tests := []struct {
		name string
		args args
		want []GroupInfo
	}{
		{"3 groups 9 competitors",
			args{competitors.NewTestCompetitors(9), calcGroupsForTest()},
			getGroupInfoWanted()},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(ToHTML(tt.args.g))
			if got := GToGroupInfo(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GToGroupInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
