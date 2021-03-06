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

func getLoc() *time.Location {
	loc, _ := time.LoadLocation("Europe/Vienna")
	return loc
}

func getCalcedPlanFor4Competitors() [][]P {
	var out [][]P
	out = append(out, []P{
		{0, 1, 1, 1, 1, time.Date(2018, 1, 20, 14, 0, 0, 0, getLoc()), -1},
		{2, 3, 1, 2, 1, time.Date(2018, 1, 20, 14, 0, 0, 0, getLoc()), -1},
	})
	out = append(out, []P{
		{0, 2, 2, 3, 2, time.Date(2018, 1, 20, 14, 15, 0, 0, getLoc()), -1},
		{1, 3, 2, 4, 2, time.Date(2018, 1, 20, 14, 15, 0, 0, getLoc()), -1},
	})
	out = append(out, []P{
		{0, 3, 3, 5, 3, time.Date(2018, 1, 20, 14, 30, 0, 0, getLoc()), -1},
		{1, 2, 3, 6, 3, time.Date(2018, 1, 20, 14, 30, 0, 0, getLoc()), -1},
	})
	return out
}
func getGamePlanForFirstTest() GamePlan {
	var out GamePlan
	out.PairingInfo = []PairingInfo{
		{"14:00", "1", "1", "Benni", "Dani"},
		{"14:00", "2", "1", "Mona", "Andrea"},
		{"14:15", "1", "2", "Benni", "Mona"},
		{"14:15", "2", "2", "Dani", "Andrea"},
		{"14:30", "1", "3", "Benni", "Andrea"},
		{"14:30", "2", "3", "Dani", "Mona"},
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

func pairingsForAllPair() []P {
	outP := []P{
		{1, 2, -8, -1, 0, time.Date(2018, 1, 20, 14, 0, 0, 0, getLoc()), 1},
		{3, 4, -8, -2, 0, time.Date(2018, 1, 20, 14, 0, 0, 0, getLoc()), 2},
		{5, 6, -8, -3, 0, time.Date(2018, 1, 20, 14, 15, 0, 0, getLoc()), 1},
		{7, 8, -8, -4, 0, time.Date(2018, 1, 20, 14, 15, 0, 0, getLoc()), 2},
	}
	return outP
}
func getGamePlanForAllPair() GamePlan {
	var out GamePlan
	out.PairingInfo = []PairingInfo{
		{"14:00", "1", "1/8 F.", "Dani", "Mona"},
		{"14:00", "2", "1/8 F.", "Andrea", "Zoé"},
		{"14:15", "1", "1/8 F.", "Andreas", "Bernhard"},
		{"14:15", "2", "1/8 F.", "Florian", "Simon"},
	}
	return out
}
func TestAllPairsToGamePlan(t *testing.T) {
	type args struct {
		c  []competitors.C
		ap []P
	}
	tests := []struct {
		name string
		args args
		want GamePlan
	}{
		{"4 competitors", args{competitors.NewTestCompetitors(9), pairingsForAllPair()}, getGamePlanForAllPair()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllPairsToGamePlan(tt.args.c, tt.args.ap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPairsToGamePlan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundToInfo(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{"Round 4", 4, "4"},
		{"Round 0", 0, ""},
		{"Round -1", -1, "Finale"},
		{"Round -8", -8, "1/8 F."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundToInfo(tt.args); got != tt.want {
				t.Errorf("roundToInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
