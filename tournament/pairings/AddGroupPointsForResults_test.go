package pairings

import (
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

type args struct {
	pairings []Pairing
	results  Results
}

func getArgsFor5() args {
	var out args
	out.pairings = append(out.pairings, Pairing{0, 1, 1, 1, 1})
	out.pairings = append(out.pairings, Pairing{2, 3, 2, 2, 1})
	out.pairings = append(out.pairings, Pairing{0, 4, 3, 3, 1})
	out.pairings = append(out.pairings, Pairing{1, 3, 4, 4, 1})
	out.results = make(map[int]Result)
	out.results[1] = Result{5, 5}
	out.results[2] = Result{3, 1}
	out.results[3] = Result{1, 2}
	out.results[4] = Result{4, 3}
	return out
}
func getArgsFor2() args {
	var out args
	out.pairings = append(out.pairings, Pairing{0, 1, 1, 1, 1})
	out.pairings = append(out.pairings, Pairing{0, 1, 2, 2, 1})
	out.pairings = append(out.pairings, Pairing{0, 1, 3, 3, 1})
	out.pairings = append(out.pairings, Pairing{0, 1, 4, 4, 1})
	out.results = make(map[int]Result)
	out.results[1] = Result{1, 5}
	out.results[2] = Result{3, 1}
	out.results[3] = Result{4, 5}
	out.results[4] = Result{4, 4}
	return out
}
func getArgsForError() args {
	var out args
	out.pairings = append(out.pairings, Pairing{0, 5, 1, 1, 1})
	out.pairings = append(out.pairings, Pairing{0, 5, 2, 2, 1})
	out.results = make(map[int]Result)
	out.results[1] = Result{1, 5}
	out.results[5] = Result{4, 4}
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
	}{
		{name: "Two competitors 4 games", args: getArgsFor2(), wantErr: false, resultPoints: []int{4, 7}},
		{name: "Five competitors 4 games", args: getArgsFor5(), wantErr: false, resultPoints: []int{1, 4, 3, 0, 3}},
		{name: "Error wanted because result found of id where is no pairing", args: getArgsForError(), wantErr: true, resultPoints: []int{}},
	}
	for _, tt := range tests {
		competitors.ClearPoints(cs)
		t.Run(tt.name, func(t *testing.T) {
			if err := AddGroupPointsForResults(cs, tt.args.pairings, tt.args.results, calc); (err != nil) != tt.wantErr {
				t.Errorf("AddGroupPointsForResults() error = %v, wantErr %v", err, tt.wantErr)

			}
			for i, r := range tt.resultPoints {

				if competitors.GetCompetitor(cs, i).GetPoints() != r {
					t.Errorf("competitorPoints %v, wanted %v", competitors.GetCompetitor(cs, i).GetPoints(), r)

				}
			}
		})
	}
}
