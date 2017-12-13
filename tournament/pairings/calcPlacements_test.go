package pairings

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

var testCompetitors competitors.Competitors

func generatePairings3() []Pairing {
	out, _ := CalcPairings(testCompetitors.Items, 1)
	for i := range out {
		out[i].ID = i
	}

	return out
}
func generatePairingResults3() []PairingResult {
	out := make([]PairingResult, 3)
	out[0] = PairingResult{0, 3, 2}
	out[1] = PairingResult{1, 4, 2}
	out[2] = PairingResult{2, 0, 2}

	return out
}
func generateExpected3() []placement {
	out := make([]placement, 3)

	testCompetitors.Items[0].GroupPoints = 6
	out[0] = placement{testCompetitors.Items[0], 1}
	testCompetitors.Items[2].GroupPoints = 3
	out[1] = placement{testCompetitors.Items[2], 2}
	testCompetitors.Items[1].GroupPoints = 0
	out[2] = placement{testCompetitors.Items[1], 3}
	return out
}

func Test_calcGroupPlacements_calcPlacements(t *testing.T) {
	testCompetitors.Items = testCompetitors.Items[:0]
	testCompetitors = competitors.NewTestCompetitors(3)
	type input struct {
		pairings      []Pairing
		pairingResult []PairingResult
	}
	tests := []struct {
		name    string
		fields  input
		want    []placement
		wantErr bool
	}{
		{name: "3 Competitors", fields: input{generatePairings3(), generatePairingResults3()}, want: generateExpected3(), wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cgp := calcGroupPlacements{
				pairings:      tt.fields.pairings,
				pairingResult: tt.fields.pairingResult,
			}
			got, err := cgp.calcPlacements(tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0))
			if (err != nil) != tt.wantErr {
				t.Errorf("calcGroupPlacements.calcPlacements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcGroupPlacements.calcPlacements() = %v, want %v", got, tt.want)
			}
		})
	}
}
