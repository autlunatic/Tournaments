package groups

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func calcGroupsForTest() []G {
	c := competitors.NewTestCompetitors(9)
	g, _ := CalcGroups(c, 3)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GToGroupInfo(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GToGroupInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
