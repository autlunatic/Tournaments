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
	c[0].AddResult(competitors.ResultPoints{5, 0, 2, 0, -1})
	c[1].AddResult(competitors.ResultPoints{2, 0, 1, 0, -1})
	c[2].AddResult(competitors.ResultPoints{4, 0, 0, 0, -1})
	g, _ := CalcGroups(c, 3)
	return g[0]
}
func getGroupInfoWanted2() GroupInfo {
	out := GroupInfo{1, []CompetitorInfos{

		CompetitorInfos{"Mona", 7, 0, 3, 0},
		CompetitorInfos{"Benni", 6, 0, 3, 0},
		CompetitorInfos{"Dani", 4, 0, 3, 0},
	}}

	return out
}
func getGroupInfoWanted() GroupInfo {
	out := GroupInfo{1, []CompetitorInfos{
		CompetitorInfos{"Mona", 3, 0, 3, 0},
		CompetitorInfos{"Dani", 2, 0, 2, 0},
		CompetitorInfos{"Benni", 1, 0, 1, 0},
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
