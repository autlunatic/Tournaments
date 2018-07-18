package pairings

import (
	"reflect"
	"testing"
	"time"

	"github.com/autlunatic/Tournaments/tournament/detail"
)

func detailsWithoutTime() detail.D {
	return detail.D{NumberOfParallelGames: 3, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 5}
}
func detailsWithTime() detail.D {
	return detail.D{NumberOfParallelGames: 3, MinutesPerGame: 5, MinutesAvailForGroupsPhase: 5, FinalsStartTime: time.Date(2018, 1, 20, 14, 20, 0, 0, time.UTC)}
}

func wantedPairs() []P {
	return []P{{1, 2, -8, -1, 0, time.Date(2018, 1, 20, 14, 20, 0, 0, time.UTC), -1},
		{3, 4, -8, -2, 0, time.Date(2018, 1, 20, 14, 20, 0, 0, time.UTC), -1},
		{5, 6, -8, -3, 0, time.Date(2018, 1, 20, 14, 20, 0, 0, time.UTC), -1},
		{7, 8, -8, -4, 0, time.Date(2018, 1, 20, 14, 25, 0, 0, time.UTC), -1},
		{9, 10, -8, -5, 0, time.Date(2018, 1, 20, 14, 25, 0, 0, time.UTC), -1},
		{11, 12, -8, -6, 0, time.Date(2018, 1, 20, 14, 25, 0, 0, time.UTC), -1},
		{13, 14, -8, -7, 0, time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC), -1},
		{15, 16, -8, -8, 0, time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC), -1},

		{1, 4, -4, -9, 0, time.Date(2018, 1, 20, 14, 35, 0, 0, time.UTC), -1},
		{6, 7, -4, -10, 0, time.Date(2018, 1, 20, 14, 35, 0, 0, time.UTC), -1},
		{10, 11, -4, -11, 0, time.Date(2018, 1, 20, 14, 35, 0, 0, time.UTC), -1},
		{13, 16, -4, -12, 0, time.Date(2018, 1, 20, 14, 40, 0, 0, time.UTC), -1},

		{1, 6, -2, -13, 0, time.Date(2018, 1, 20, 14, 45, 0, 0, time.UTC), -1},
		{10, 16, -2, -14, 0, time.Date(2018, 1, 20, 14, 45, 0, 0, time.UTC), -1}}
}
func Test_calcTimesForFinalPairings(t *testing.T) {
	type args struct {
		lastGameGroup time.Time
		finPairs      []P
		d             detail.D
	}
	tests := []struct {
		name string
		args args
		want []P
	}{
		{"time in Details not set, go on after last GroupMatch", args{time.Date(2018, 1, 20, 14, 15, 0, 0, time.UTC), argsFor8ToFin().pairs, detailsWithoutTime()}, wantedPairs()},
		{"time in Details set, go on after last GroupMatch", args{time.Date(2018, 1, 20, 10, 15, 0, 0, time.UTC), argsFor8ToFin().pairs, detailsWithTime()}, wantedPairs()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcTimesForFinalPairings(tt.args.lastGameGroup, tt.args.finPairs, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcTimesForFinalPairings() = %v, want %v", got, tt.want)
			}
		})
	}
}
