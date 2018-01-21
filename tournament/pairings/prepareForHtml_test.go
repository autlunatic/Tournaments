package pairings

import (
	"reflect"
	"testing"
	"time"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func getTestCompetitors() []competitors.C {
	return competitors.NewTestCompetitors(4)
}

func getCalcedPlanFor4Competitors() [][]P {
	var out [][]P
	out = append(out, []P{
		{0, 1, 1, 1, 1},
		{2, 3, 1, 2, 1},
	})
	out = append(out, []P{
		{0, 2, 1, 3, 2},
		{1, 3, 1, 4, 2},
	})
	out = append(out, []P{
		{0, 3, 1, 5, 3},
		{1, 2, 1, 6, 3},
	})
	return out
}
func getGamePlanForFirstTest() GamePlan {
	var out GamePlan
	out.pairingInfo = []PairingInfo{
		{time.Date(2018, 1, 20, 14, 0, 0, 0, time.UTC), "1", P{0, 1, 1, 1, 1}, "1", "2"},
		{time.Date(2018, 1, 20, 14, 0, 0, 0, time.UTC), "2", P{2, 3, 1, 2, 1}, "3", "4"},
		{time.Date(2018, 1, 20, 14, 15, 0, 0, time.UTC), "1", P{0, 2, 1, 3, 2}, "1", "3"},
		{time.Date(2018, 1, 20, 14, 15, 0, 0, time.UTC), "2", P{1, 3, 1, 4, 2}, "2", "4"},
		{time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC), "1", P{0, 3, 1, 5, 3}, "1", "4"},
		{time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC), "2", P{1, 2, 1, 6, 3}, "2", "3"},
	}
	return out
}
func Test_calcedPlanToGamePlan(t *testing.T) {
	type args struct {
		startTime       time.Time
		durationPerGame int
		c               []competitors.C
		cp              [][]P
	}
	tests := []struct {
		name string
		args args
		want GamePlan
	}{
		{"4 competitors", args{time.Date(2018, 1, 20, 14, 0, 0, 0, time.UTC), 15, getTestCompetitors(), getCalcedPlanFor4Competitors()}, getGamePlanForFirstTest()},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcedPlanToGamePlan(tt.args.startTime, tt.args.durationPerGame, tt.args.c, tt.args.cp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcedPlanToGamePlan() =\n got: %v,  \n want %v", got, tt.want)
			}
		})
	}
}
