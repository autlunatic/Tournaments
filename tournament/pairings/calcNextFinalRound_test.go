package pairings

import (
	"reflect"
	"testing"
	"time"

	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

type calcNextFinalRoundArgs struct {
	pairs                 []P
	results               Results
	numberOfParallelGames int
}

func getPairingsForQuarterFinals(aIdsFrom16 bool) calcNextFinalRoundArgs {
	var offset int
	if aIdsFrom16 {
		offset = 8
	}
	outP := []P{
		{1, 2, -4, -1 - offset, 0, time.Time{}, 1},
		{3, 4, -4, -2 - offset, 0, time.Time{}, 2},
		{5, 6, -4, -3 - offset, 0, time.Time{}, 1},
		{7, 8, -4, -4 - offset, 0, time.Time{}, 2},
	}
	outR := Results{
		-1 - offset: &Result{3, 1},
		-2 - offset: &Result{5, 6},
		-3 - offset: &Result{2, 1},
		-4 - offset: &Result{1, 3},
	}
	return calcNextFinalRoundArgs{outP, outR, 2}
}

func quarterToSemi() calcNextFinalRoundArgs {
	outP := []P{
		{1, 2, -4, -1, 0, time.Time{}, -1},
		{3, 4, -4, -2, 0, time.Time{}, -1},
		{5, 6, -4, -3, 0, time.Time{}, -1},
		{7, 8, -4, -4, 0, time.Time{}, -1},
	}
	outR := Results{
		-1: &Result{3, 1},
		-2: &Result{5, 6},
	}
	return calcNextFinalRoundArgs{outP, outR, 2}
}

func getPairingsForSemiFinals() calcNextFinalRoundArgs {
	outP := []P{
		{1, 2, -2, -5, 0, time.Time{}, -1},
		{3, 4, -2, -6, 0, time.Time{}, -1},
	}
	outR := Results{
		-5: &Result{3, 1},
		-6: &Result{5, 6},
	}
	return calcNextFinalRoundArgs{outP, outR, 2}

}
func quarterToSemiOnlyOneResult() calcNextFinalRoundArgs {
	outP := []P{
		{1, 2, -4, -1, 0, time.Time{}, -1},
		{3, 4, -4, -2, 0, time.Time{}, -1},
		{5, 6, -4, -3, 0, time.Time{}, -1},
		{7, 8, -4, -4, 0, time.Time{}, -1},
	}
	outR := Results{
		-1: &Result{3, 1},
	}
	return calcNextFinalRoundArgs{outP, outR, 2}
}
func getPairingsWithInvalidResults() calcNextFinalRoundArgs {
	outP := []P{
		{1, 2, -4, -1, 0, time.Time{}, -1},
		{3, 4, -4, -2, 0, time.Time{}, -1},
		{5, 6, -4, -3, 0, time.Time{}, -1},
		{7, 8, -4, -4, 0, time.Time{}, -1},
		{1, 3, -2, -5, 0, time.Time{}, -1},
	}
	outR := Results{
		-5: &Result{3, 1},
	}
	return calcNextFinalRoundArgs{outP, outR, 2}
}
func getFinalWithResult() calcNextFinalRoundArgs {
	outP := []P{
		{1, 3, -1, -3, 0, time.Time{}, -1},
	}
	outR := Results{
		-3: &Result{3, 1},
	}
	return calcNextFinalRoundArgs{outP, outR, 2}
}

func TestCalcNextFinalRound(t *testing.T) {
	tests := []struct {
		name string
		args calcNextFinalRoundArgs
		want []P
	}{
		{"quarter to semi", getPairingsForQuarterFinals(false), []P{{1, 4, -2, -5, 0, time.Time{}, 1}, {5, 8, -2, -6, 0, time.Time{}, 2}}},
		{"quarter to semi, but the ids of quarter are coming from a round of 16", getPairingsForQuarterFinals(true), []P{{1, 4, -2, -13, 0, time.Time{}, 1}, {5, 8, -2, -14, 0, time.Time{}, 2}}},
		{"semi to finals", getPairingsForSemiFinals(), []P{{1, 4, -1, -7, 0, time.Time{}, 1}}},
		{"final should always return empty", getFinalWithResult(), []P{}},
		{"quarter to semi 1 result", quarterToSemiOnlyOneResult(), []P{{1, -1, -2, -5, 0, time.Time{}, 1}}},
		{"quarter to semi 2 results", quarterToSemi(), []P{{1, 4, -2, -5, 0, time.Time{}, 1}}},
		{"quarter to semi invalid results", getPairingsWithInvalidResults(), []P{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcNextFinalRound(tt.args.pairs, tt.args.results, tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0), tt.args.numberOfParallelGames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcNextFinalRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func argsForSemi() []P {
	outP := []P{
		{1, 2, -4, -1, 0, time.Time{}, -1},
		{3, 4, -4, -2, 0, time.Time{}, -1},
		{5, 6, -4, -3, 0, time.Time{}, -1},
		{7, 8, -4, -4, 0, time.Time{}, -1},
		{1, 3, -2, -5, 0, time.Time{}, -1},
		{5, 7, -2, -6, 0, time.Time{}, -1},
	}
	return outP
}
func wantForSemi() []P {
	outP := []P{
		{1, 3, -2, -5, 0, time.Time{}, -1},
		{5, 7, -2, -6, 0, time.Time{}, -1},
	}
	return outP
}

func argsForIncompleteSemi() []P {
	outP := []P{
		{1, 2, -4, -1, 0, time.Time{}, -1},
		{5, 6, -4, -3, 0, time.Time{}, -1},
		{7, 8, -4, -4, 0, time.Time{}, -1},
		{5, 7, -2, -6, 0, time.Time{}, -1},
		{3, 4, -4, -2, 0, time.Time{}, -1},
	}
	return outP
}
func wantForIncompleteSemi() []P {
	outP := []P{
		{5, 7, -2, -6, 0, time.Time{}, -1},
	}
	return outP
}

func wantForSemiFirst() []P {
	outP := []P{
		{1, 2, -4, -1, 0, time.Time{}, -1},
		{3, 4, -4, -2, 0, time.Time{}, -1},
		{5, 6, -4, -3, 0, time.Time{}, -1},
		{7, 8, -4, -4, 0, time.Time{}, -1},
	}
	return outP

}
func Test_filterOutFirstFinalRound(t *testing.T) {
	tests := []struct {
		name     string
		finPairs []P
		want     []P
	}{
		{"quarter and semi finals only semi are returned", argsForSemi(), wantForSemiFirst()},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterOutFirstFinalRound(tt.finPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOutFirstFinalRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
func getArgsForRecalc(resultForPID0 bool) calcNextFinalRoundArgs {
	outP := []P{
		{1, 2, -4, -1, 0, time.Time{}, 1},
		{3, 4, -4, -2, 0, time.Time{}, 2},
		{5, 6, -4, -3, 0, time.Time{}, 3},
		{7, 8, -4, -4, 0, time.Time{}, 4},
	}
	if !resultForPID0 {
		outP = append(outP, P{1, 3, -2, -5, 0, time.Time{}, -1})
	}
	outR := Results{
		-1: &Result{3, 1},
		-2: &Result{3, 1},
		-3: &Result{3, 1},
		-4: &Result{3, 1},
	}
	if resultForPID0 {
		outR[0] = &Result{3, 1}
	} else {
		outR[-5] = &Result{3, 1}
	}
	return calcNextFinalRoundArgs{outP, outR, 4}
}
func wantedRecalc(resultForPID0 bool) []P {
	out := []P{
		{1, 2, -4, -1, 0, time.Time{}, 1},
		{3, 4, -4, -2, 0, time.Time{}, 2},
		{5, 6, -4, -3, 0, time.Time{}, 3},
		{7, 8, -4, -4, 0, time.Time{}, 4},
		{1, 3, -2, -5, 0, time.Time{}, 1},
		{5, 7, -2, -6, 0, time.Time{}, 2},
	}
	if !resultForPID0 {
		out = append(out, P{1, -1, -1, -7, 0, time.Time{}, 3})
	}
	return out
}
func argsFor8ToFin() calcNextFinalRoundArgs {
	outP := []P{
		{1, 2, -8, -1, 0, time.Time{}, 1},
		{3, 4, -8, -2, 0, time.Time{}, 2},
		{5, 6, -8, -3, 0, time.Time{}, 3},
		{7, 8, -8, -4, 0, time.Time{}, 1},
		{9, 10, -8, -5, 0, time.Time{}, 2},
		{11, 12, -8, -6, 0, time.Time{}, 3},
		{13, 14, -8, -7, 0, time.Time{}, 1},
		{15, 16, -8, -8, 0, time.Time{}, 2},

		{1, 4, -4, -9, 0, time.Time{}, 3},
		{6, 7, -4, -10, 0, time.Time{}, 1},
		{10, 11, -4, -11, 0, time.Time{}, 2},
		{13, 16, -4, -12, 0, time.Time{}, 3},

		{1, 6, -2, -13, 0, time.Time{}, 1},
		{10, 16, -2, -14, 0, time.Time{}, 2},
	}
	outR := Results{
		-1:  &Result{3, 1},
		-2:  &Result{3, 4},
		-3:  &Result{3, 5},
		-4:  &Result{2, 1},
		-5:  &Result{3, 9},
		-6:  &Result{3, 1},
		-7:  &Result{3, 1},
		-8:  &Result{1, 2},
		-9:  &Result{3, 1},
		-10: &Result{3, 1},
		-11: &Result{3, 1},
		-12: &Result{3, 6},
		-13: &Result{3, 1},
		-14: &Result{3, 1},
	}
	return calcNextFinalRoundArgs{outP, outR, 3}
}

func wantFor8ToFin() []P {
	out := argsFor8ToFin().pairs
	out = append(out, P{1, 10, -1, -15, 0, time.Time{}, 3})
	return out
}

func TestRecalcFinals(t *testing.T) {
	tests := []struct {
		name string
		args calcNextFinalRoundArgs
		want []P
	}{
		{"quarter to semi", getArgsForRecalc(false), wantedRecalc(false)},
		{"roundOfEight to final", argsFor8ToFin(), wantFor8ToFin()},
		{"quarter to semi with result for 0 PairingID", getArgsForRecalc(true), wantedRecalc(true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecalcFinals(tt.args.pairs, tt.args.results, tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0), tt.args.numberOfParallelGames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecalcFinals() = %v, want %v", got, tt.want)
			}
		})
	}
}
