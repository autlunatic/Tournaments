package pairings

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

type testRenderResultArgs struct {
	c   []competitors.C
	p   []P
	r   Results
	tpc tournamentPoints.TournamentPointCalcer
}

func getRenderArgs() testRenderResultArgs {
	var out testRenderResultArgs
	a := getArgsFor5()
	out.p = a.pairings
	out.r = a.results
	out.c = competitors.NewTestCompetitors(5)
	out.tpc = tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0)
	return out

}
func getRenderWanted() []ResultInfo {

	out := []ResultInfo{

		{Comp1Name: "Benni", Comp2Name: "Dani", Group1Pts: 1, Group2Pts: 1, Pairing1Pts: 5, Pairing2Pts: 5, PairingID: 1},
		{Comp1Name: "Mona", Comp2Name: "Andrea", Group1Pts: 3, Group2Pts: 0, Pairing1Pts: 3, Pairing2Pts: 1, PairingID: 2},
		{Comp1Name: "Benni", Comp2Name: "Zoé", Group1Pts: 0, Group2Pts: 3, Pairing1Pts: 1, Pairing2Pts: 2, PairingID: 3},
		{Comp1Name: "Dani", Comp2Name: "Andrea", Group1Pts: 3, Group2Pts: 0, Pairing1Pts: 4, Pairing2Pts: 3, PairingID: 4},
	}
	return out
}

func TestResultsToResultInfo(t *testing.T) {
	tests := []struct {
		name string
		args testRenderResultArgs
		want []ResultInfo
	}{
		// TODO: Add test cases.
		{"5 competitors 4 results", getRenderArgs(), getRenderWanted()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResultsToResultInfo(tt.args.c, tt.args.p, tt.args.r, tt.args.tpc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResultsToResultInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
