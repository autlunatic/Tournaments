package tournamentPoints

import (
	"testing"
)

func TestSoccer(t *testing.T) {

}

func TestSimpleTournamentPointCalc_calc(t *testing.T) {
	type args struct {
		gamePoints1 int
		GamePoints2 int
	}
	calc := NewSimpleTournamentPointCalc(1, 3, 0)
	tests := []struct {
		name                  string
		s                     SimpleTournamentPointCalc
		args                  args
		wantTournamentPoints1 int
		wantTournamentPoints2 int
	}{
		{"simplesoccer Draw", calc, args{0, 0}, 1, 1},
		{"simplesoccer team 1 win", calc, args{1, 0}, 3, 0},
		{"simplesoccer team 2 win", calc, args{1, 3}, 0, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTournamentPoints1, gotTournamentPoints2 := tt.s.Calc(tt.args.gamePoints1, tt.args.GamePoints2)
			if gotTournamentPoints1 != tt.wantTournamentPoints1 {
				t.Errorf("SimpleTournamentPointCalc.calc() gotTournamentPoints1 = %v, want %v", gotTournamentPoints1, tt.wantTournamentPoints1)
			}
			if gotTournamentPoints2 != tt.wantTournamentPoints2 {
				t.Errorf("SimpleTournamentPointCalc.calc() gotTournamentPoints2 = %v, want %v", gotTournamentPoints2, tt.wantTournamentPoints2)
			}
		})
	}
}
