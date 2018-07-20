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
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 1, Comp2Name: "Dani", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 0, PairingInfo: "0", Done: false},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 2, Comp2Name: "Mona", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 4, PairingInfo: "4", Done: false},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 3, Comp2Name: "Andrea", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 8, PairingInfo: "8", Done: false},
		{Comp1ID: 0, Comp1Name: "Benni", Comp2ID: 4, Comp2Name: "Zoé", Group1Pts: 0, Group2Pts: 0, Pairing1Pts: 0, Pairing2Pts: 0, PairingID: 12, PairingInfo: "12", Done: false},
	}

	return out
}

func getWantedCompPageInfo() CompetitorPageInfo {
	ttp := ttForCompPage.Pairings
	fmt.Println(ttp)
	p := []pairings.P{ttp[0],
		ttp[4],
		ttp[8],
		ttp[12],
	}

	fmt.Println(p)
	fmt.Println(ttForCompPage.Competitors)
	out := CompetitorPageInfo{g: ttForCompPage.Groups[0],
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
