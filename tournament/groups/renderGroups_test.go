package groups

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func calcGroupsForTest() G {
	c := competitors.NewTestCompetitors(9)
	g, _ := CalcGroups(c, 3)
	return g[0]
}

func getGroupInfoWanted() GroupInfo {
	out := GroupInfo{1, []CompetitorInfos{CompetitorInfos{"Benni", 0, 1},
		CompetitorInfos{"Dani", 0, 2},
		CompetitorInfos{"Mona", 0, 3},
	}}

	return out
}
func TestGToGroupInfo(t *testing.T) {
	type args struct {
		c []competitors.C
		g G
	}
	tests := []struct {
		name string
		args args
		want GroupInfo
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
