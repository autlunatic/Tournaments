package tournament

import (
	"reflect"
	"testing"
	"time"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

func (t *T) AddCompetitor(c competitors.C) {
	t.Competitors = append(t.Competitors, c)
}

func NewTestTournament() T {
	td := detail.D{
		MinutesAvailForGroupsPhase: 90,
		MinutesPerGame:             15,
		NumberOfParallelGames:      4}
	t := T{Details: td}
	t.Competitors = competitors.NewTestCompetitors(9)

	t.FinalPairings = []pairings.P{
		pairings.P{Competitor1ID: 1, Competitor2ID: 2, Round: -4, ID: -1, GroupID: 0, StartTime: time.Time{}, Court: 2},
		pairings.P{Competitor1ID: 3, Competitor2ID: 4, Round: -4, ID: -2, GroupID: 0, StartTime: time.Time{}, Court: 2},
		pairings.P{Competitor1ID: 5, Competitor2ID: 6, Round: -4, ID: -3, GroupID: 0, StartTime: time.Time{}, Court: 1},
		pairings.P{Competitor1ID: 7, Competitor2ID: 8, Round: -4, ID: -4, GroupID: 0, StartTime: time.Time{}, Court: 2},
		pairings.P{Competitor1ID: 1, Competitor2ID: 3, Round: -2, ID: -5, GroupID: 0, StartTime: time.Time{}, Court: 1},
		pairings.P{Competitor1ID: 5, Competitor2ID: 7, Round: -2, ID: -6, GroupID: 0, StartTime: time.Time{}, Court: 2},
		pairings.P{Competitor1ID: 1, Competitor2ID: 0, Round: -1, ID: -7, GroupID: 0, StartTime: time.Time{}, Court: 1},
	}

	t.Build()

	return t
}

func TestAddCompetitor(t *testing.T) {
	tournament := NewTestTournament()
	tournament.AddCompetitor(competitors.New("NotInTestList", 0))
	if len(tournament.Competitors) != 10 {
		t.Error("Should be ten competitors")
	}
	if tournament.Competitors[0].Name() != "Benni" {
		t.Error("first competitors name should be Benni")
	}

}

var trmntForTestGetPairingByID = NewTestTournament()

func TestT_GetPairingByID(t *testing.T) {
	tests := []struct {
		name    string
		t       T
		ID      int
		want    pairings.P
		wantErr bool
	}{
		{"empty Tournament returns error", T{}, 27, pairings.P{}, true},
		{"TestTournament with correct ID returns P", trmntForTestGetPairingByID, 2, trmntForTestGetPairingByID.Pairings[1], false},
		{"TestTournament with Finalpairings should return Final", trmntForTestGetPairingByID, -1, trmntForTestGetPairingByID.FinalPairings[0], false},
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
