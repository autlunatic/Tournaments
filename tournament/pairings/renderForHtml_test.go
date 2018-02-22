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
		{0, 1, 1, 1, 1, time.Date(2018, 1, 20, 14, 0, 0, 0, time.UTC)},
		{2, 3, 1, 2, 1, time.Date(2018, 1, 20, 14, 0, 0, 0, time.UTC)},
	})
	out = append(out, []P{
		{0, 2, 2, 3, 1, time.Date(2018, 1, 20, 14, 15, 0, 0, time.UTC)},
		{1, 3, 2, 4, 1, time.Date(2018, 1, 20, 14, 15, 0, 0, time.UTC)},
	})
	out = append(out, []P{
		{0, 3, 3, 5, 1, time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC)},
		{1, 2, 3, 6, 1, time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC)},
	})
	return out
}
func getGamePlanForFirstTest() GamePlan {
	var out GamePlan
	out.PairingInfo = []PairingInfo{
		{"14:00", "1", P{0, 1, 1, 1, 1, time.Date(2018, 1, 20, 14, 0, 0, 0, time.UTC)}, "Benni", "Dani"},
		{"14:00", "2", P{2, 3, 1, 2, 1, time.Date(2018, 1, 20, 14, 0, 0, 0, time.UTC)}, "Mona", "Andrea"},
		{"14:15", "1", P{0, 2, 2, 3, 1, time.Date(2018, 1, 20, 14, 15, 0, 0, time.UTC)}, "Benni", "Mona"},
		{"14:15", "2", P{1, 3, 2, 4, 1, time.Date(2018, 1, 20, 14, 15, 0, 0, time.UTC)}, "Dani", "Andrea"},
		{"14:30", "1", P{0, 3, 3, 5, 1, time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC)}, "Benni", "Andrea"},
		{"14:30", "2", P{1, 2, 3, 6, 1, time.Date(2018, 1, 20, 14, 30, 0, 0, time.UTC)}, "Dani", "Mona"},
	}
	return out
}
func Test_calcedPlanToGamePlan(t *testing.T) {
	type args struct {
		c  []competitors.C
		cp [][]P
	}
	tests := []struct {
		name string
		args args
		want GamePlan
	}{
		{"4 competitors", args{getTestCompetitors(), getCalcedPlanFor4Competitors()}, getGamePlanForFirstTest()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcedPlanToGamePlan(tt.args.c, tt.args.cp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcedPlanToGamePlan() =\n got: %v,  \n want %v", got, tt.want)
			}
		})
	}
}
