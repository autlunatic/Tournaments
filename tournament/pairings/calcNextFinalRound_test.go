package pairings

import (
	"reflect"
	"testing"
	"time"

	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

type calcNextFinalRoundArgs struct {
	pairs   []P
	results Results
}

func getPairingsForQuarterFinals(aIdsFrom16 bool) calcNextFinalRoundArgs {
	var offset int
	if aIdsFrom16 {
		offset = 8
	}
	outP := []P{
		{1, 2, -4, -1 - offset, 0, time.Time{}},
		{3, 4, -4, -2 - offset, 0, time.Time{}},
		{5, 6, -4, -3 - offset, 0, time.Time{}},
		{7, 8, -4, -4 - offset, 0, time.Time{}},
	}
	outR := Results{
		-1 - offset: &Result{3, 1},
		-2 - offset: &Result{5, 6},
		-3 - offset: &Result{2, 1},
		-4 - offset: &Result{1, 3},
	}
	return calcNextFinalRoundArgs{outP, outR}
}
func getPairingsForSemiFinals() calcNextFinalRoundArgs {
	outP := []P{
		{1, 2, -2, -5, 0, time.Time{}},
		{3, 4, -2, -6, 0, time.Time{}},
	}
	outR := Results{
		-5: &Result{3, 1},
		-6: &Result{5, 6},
	}
	return calcNextFinalRoundArgs{outP, outR}
}

func TestCalcNextFinalRound(t *testing.T) {
	tests := []struct {
		name string
		args calcNextFinalRoundArgs
		want []P
	}{
		{"quarter to semi", getPairingsForQuarterFinals(false), []P{{1, 4, -2, -5, 0, time.Time{}}, {5, 8, -2, -6, 0, time.Time{}}}},
		{"quarter to semi, but the ids of quarter are coming from a round of 16", getPairingsForQuarterFinals(true), []P{{1, 4, -2, -13, 0, time.Time{}}, {5, 8, -2, -14, 0, time.Time{}}}},
		{"semi to finals", getPairingsForSemiFinals(), []P{{1, 4, -1, -7, 0, time.Time{}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcNextFinalRound(tt.args.pairs, tt.args.results, tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcNextFinalRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
