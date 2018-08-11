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
func calcGroupsForTest2() G {
	c := competitors.NewTestCompetitors(9)
	g, _ := CalcGroups(c, 3)
	return g[0]
}
func getGroupInfoWanted2() GroupInfo {
	out := GroupInfo{1, []CompetitorInfos{

		CompetitorInfos{"Benni", 5, 3, 5, 0},
		CompetitorInfos{"Mona", 4, 3, 4, 3},
		CompetitorInfos{"Dani", 2, 3, 2, 3},
	}}

	return out
}
func getGroupInfoWanted() GroupInfo {
	out := GroupInfo{1, []CompetitorInfos{
		CompetitorInfos{"Mona", 0, 3, 0, 3},
		CompetitorInfos{"Dani", 0, 2, 0, 2},
		CompetitorInfos{"Benni", 0, 1, 0, 1},
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
		{"3 groups 9 competitors", args{competitors.NewTestCompetitors(9), calcGroupsForTest()}, getGroupInfoWanted()},
		{"3 groups 9 competitors all same GroupPoints", args{competitors.NewTestCompetitors(9), calcGroupsForTest2()}, getGroupInfoWanted2()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GToGroupInfo(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GToGroupInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
