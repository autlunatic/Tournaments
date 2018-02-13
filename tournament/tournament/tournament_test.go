package tournament

import (
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

func (t *T) AddCompetitor(c competitors.C) {
	t.Competitors = append(t.Competitors, c)
}

func newTestTournament() T {
	td := &detail.D{
		MinutesAvailForGroupsPhase: 90,
		MinutesPerGame:             15,
		NumberOfParallelGames:      4}
	tournament := NewTournament(*td)
	tournament.Competitors = competitors.NewTestCompetitors(9)
	tournament.Build()

	return tournament
}

func TestAddCompetitor(t *testing.T) {
	tournament := newTestTournament()
	tournament.AddCompetitor(competitors.New("NotInTestList", 0))
	if len(tournament.Competitors) != 10 {
		t.Error("Should be ten competitors")
	}
	if tournament.Competitors[0].Name() != "Benni" {
		t.Error("first competitors name should be Benni")
	}

}

var trmntForTestGetPairingByID = newTestTournament()

func TestT_GetPairingByID(t *testing.T) {
	tests := []struct {
		name    string
		t       T
		ID      int
		want    pairings.P
		wantErr bool
	}{
		{"empty Tournament returns error", T{}, 27, pairings.P{}, true},
		{"TestTournament with correct ID returns P", trmntForTestGetPairingByID, 2, trmntForTestGetPairingByID.Pairings[2], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.GetPairingByID(tt.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("T.GetPairingByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("T.GetPairingByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
