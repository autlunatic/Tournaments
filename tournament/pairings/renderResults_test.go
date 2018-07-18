package pairings

import (
	"reflect"
	"testing"
	"time"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
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
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 1, Comp2Name: "Dani", Group1Pts: 1, Group2Pts: 1, Pairing1Pts: 5, Pairing2Pts: 5, PairingID: 1, PairingInfo: "1", Done: true},
		{Comp1ID: 2, Comp1Name: "Mona", Comp2ID: 3, Comp2Name: "Andrea", Group1Pts: 3, Group2Pts: 0, Pairing1Pts: 3, Pairing2Pts: 1, PairingID: 2, PairingInfo: "2", Done: true},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 4, Comp2Name: "Zoé", Group1Pts: 0, Group2Pts: 3, Pairing1Pts: 1, Pairing2Pts: 2, PairingID: 3, PairingInfo: "3", Done: true},
		{Comp1ID: 1, Comp1Name: "Dani", Comp2ID: 3, Comp2Name: "Andrea", Group1Pts: 3, Group2Pts: 0, Pairing1Pts: 4, Pairing2Pts: 3, PairingID: 4, PairingInfo: "4", Done: true},
	}
	return out
}

func getFilteredWanted() []ResultInfo {
	out := []ResultInfo{
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 1, Comp2Name: "Dani", Group1Pts: 1, Group2Pts: 1, Pairing1Pts: 5, Pairing2Pts: 5, PairingID: 1, PairingInfo: "1", Done: true},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 4, Comp2Name: "Zoé", Group1Pts: 0, Group2Pts: 3, Pairing1Pts: 1, Pairing2Pts: 2, PairingID: 3, PairingInfo: "3", Done: true},
	}
	return out
}

func getRenderArgsNoResults() testRenderResultArgs {
	var out testRenderResultArgs
	a := getArgsFor5()
	out.p = a.pairings
	out.c = competitors.NewTestCompetitors(5)
	out.tpc = tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0)
	return out

}
func getRenderWanted2() []ResultInfo {
	out := []ResultInfo{
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 1, Comp2Name: "Dani", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 1, PairingInfo: "1", Done: false},
		{Comp1ID: 2, Comp1Name: "Mona", Comp2ID: 3, Comp2Name: "Andrea", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 2, PairingInfo: "2", Done: false},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 4, Comp2Name: "Zoé", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 3, PairingInfo: "3", Done: false},
		{Comp1ID: 1, Comp1Name: "Dani", Comp2ID: 3, Comp2Name: "Andrea", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 4, PairingInfo: "4", Done: false},
	}
	return out
}

func testArgsFinals() testRenderResultArgs {
	var out testRenderResultArgs
	out.c = competitors.NewTestCompetitors(5)
	out.tpc = tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0)
	var p []P
	p = append(p, P{0, 1, -4, -1, 1, time.Time{}, -1})
	p = append(p, P{2, 3, -4, -2, 1, time.Time{}, -1})
	p = append(p, P{0, 4, -2, -3, 1, time.Time{}, -1})
	p = append(p, P{1, 3, -2, -4, 1, time.Time{}, -1})
	out.p = p
	out.r = make(map[int]*Result)
	out.r[-1] = &Result{5, 5}
	out.r[-2] = &Result{3, 1}
	out.r[-3] = &Result{1, 2}
	return out

}
func getRenderWantedFinals() []ResultInfo {
	out := []ResultInfo{
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 1, Comp2Name: "Dani", Group1Pts: 1, Group2Pts: 1, Pairing1Pts: 5, Pairing2Pts: 5, PairingID: -1, PairingInfo: "1/4 F.", Done: true},
		{Comp1ID: 2, Comp1Name: "Mona", Comp2ID: 3, Comp2Name: "Andrea", Group1Pts: 3, Group2Pts: 0, Pairing1Pts: 3, Pairing2Pts: 1, PairingID: -2, PairingInfo: "1/4 F.", Done: true},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 4, Comp2Name: "Zoé", Group1Pts: 0, Group2Pts: 3, Pairing1Pts: 1, Pairing2Pts: 2, PairingID: -3, PairingInfo: "1/2 F.", Done: true},
		{Comp1ID: 1, Comp1Name: "Dani", Comp2ID: 3, Comp2Name: "Andrea", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: -4, PairingInfo: "1/2 F.", Done: false},
	}
	return out
}
func TestResultsToResultInfo(t *testing.T) {
	tests := []struct {
		name string
		args testRenderResultArgs
		want []ResultInfo
	}{
		{"5 competitors 4 results", getRenderArgs(), getRenderWanted()},
		{"5 Competitors no results", getRenderArgsNoResults(), getRenderWanted2()},
		{"5 finalPairings with results", testArgsFinals(), getRenderWantedFinals()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResultsToResultInfo(tt.args.c, tt.args.p, tt.args.r, tt.args.tpc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResultsToResultInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterResultInfoByCompID(t *testing.T) {
	type args struct {
		ris    []ResultInfo
		compID int
	}
	tests := []struct {
		name string
		args args
		want []ResultInfo
	}{
		{"5 Competitors 4 results", args{getRenderWanted(), 0}, getFilteredWanted()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterResultInfoByCompID(tt.args.ris, tt.args.compID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterResultInfoByCompID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairingIDToInfo(t *testing.T) {
	type args struct {
		id    int
		round int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"positive ID", args{1, 1}, "1"},
		{"Negative ID", args{-1, -2}, "1/2 F."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairingIDToInfo(tt.args.id, tt.args.round); got != tt.want {
				t.Errorf("pairingIDToInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
