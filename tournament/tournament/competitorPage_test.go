package tournament

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/autlunatic/Tournaments/tournament/pairings"
)

var ttForCompPage = NewTestTournament()

func getResultInfoWanted() []pairings.ResultInfo {
	out := []pairings.ResultInfo{
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 1, Comp2Name: "Dani", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 1, Done: false},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 4, Comp2Name: "Zoé", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 3, Done: false},
	}
	return out
}

func getWantedCompPageInfo() CompetitorPageInfo {
	fmt.Println(ttForCompPage.Pairings)
	ttp := ttForCompPage.Pairings
	p := []pairings.P{ttp[1],
		ttp[5],
		ttp[8],
		ttp[13],
	}

	out := CompetitorPageInfo{g: ttForCompPage.Groups[1],
		pairs: p,
		ri:    getResultInfoWanted(),
	}

	return out
}

func TestToCompetitorPageInfo(t *testing.T) {
	type args struct {
		competitorID int
		t            T
	}
	tests := []struct {
		name string
		args args
		want CompetitorPageInfo
	}{
		{"TestTournament Bennis Page", args{0, ttForCompPage}, getWantedCompPageInfo()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCompetitorPageInfo(tt.args.competitorID, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToCompetitorPageInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
