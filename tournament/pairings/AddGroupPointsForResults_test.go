package pairings

import (
	"testing"
	"time"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

type args struct {
	pairings []P
	results  Results
}

func getArgsFor5() args {
	var out args
	out.pairings = append(out.pairings, P{0, 1, 1, 1, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{2, 3, 2, 2, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{0, 4, 3, 3, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{1, 3, 4, 4, 1, time.Time{}, -1})
	out.results = make(map[int]*Result)
	out.results[1] = &Result{5, 5}
	out.results[2] = &Result{3, 1}
	out.results[3] = &Result{1, 2}
	out.results[4] = &Result{4, 3}
	return out
}
func getArgsFor2() args {
	var out args
	out.pairings = append(out.pairings, P{0, 1, 1, 1, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{0, 1, 2, 2, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{0, 1, 3, 3, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{0, 1, 4, 4, 1, time.Time{}, -1})
	out.results = make(map[int]*Result)
	out.results[1] = &Result{1, 5}
	out.results[2] = &Result{3, 1}
	out.results[3] = &Result{4, 5}
	out.results[4] = &Result{4, 4}
	return out
}
func getArgsForError1() args {
	var out args
	out.pairings = append(out.pairings, P{0, 5, 1, 1, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{0, 5, 2, 2, 1, time.Time{}, -1})
	out.results = make(map[int]*Result)
	out.results[1] = &Result{1, 5}
	out.results[5] = &Result{4, 4}
	return out
}
func getArgsForError2() args {
	var out args
	out.pairings = append(out.pairings, P{5, 0, 1, 1, 1, time.Time{}, -1})
	out.pairings = append(out.pairings, P{5, 0, 2, 2, 1, time.Time{}, -1})
	out.results = make(map[int]*Result)
	out.results[1] = &Result{1, 5}
	out.results[5] = &Result{4, 4}
	return out
}

func TestAddGroupPointsForResults(t *testing.T) {
	cs := competitors.NewTestCompetitors(5)
	calc := tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0)
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		resultPoints []int
		gamePoints   []int
	}{
		{name: "Two competitors 4 games", args: getArgsFor2(), wantErr: false, resultPoints: []int{4, 7}, gamePoints: []int{12, 15}},
		{name: "Five competitors 4 games", args: getArgsFor5(), wantErr: false, resultPoints: []int{1, 4, 3, 0, 3}, gamePoints: []int{6, 9, 3, 4, 2}},
		{name: "Error wanted because Competitor 1 was not found", args: getArgsForError1(), wantErr: true, resultPoints: []int{}},
		{name: "Error wanted because Competitor 2 was not found", args: getArgsForError2(), wantErr: true, resultPoints: []int{}},
	}
	for _, tt := range tests {
		competitors.ClearPoints(cs)
		t.Run(tt.name, func(t *testing.T) {
			if err := AddPointsForResults(cs, tt.args.pairings, tt.args.results, calc); (err != nil) != tt.wantErr {
				t.Errorf("AddGroupPointsForResults() error = %v, wantErr %v", err, tt.wantErr)

			}
			for i, r := range tt.resultPoints {
				if competitors.GetCompetitor(cs, i).GetPoints() != r {
					t.Errorf("competitorPoints %v, wanted %v", competitors.GetCompetitor(cs, i).GetPoints(), r)
				}
			}
			for i, r := range tt.gamePoints {
				if competitors.GetCompetitor(cs, i).GetGamePoints() != r {
					t.Errorf("gamePoints %v, wanted %v", competitors.GetCompetitor(cs, i).GetGamePoints(), r)
				}
			}

		})
	}
}
