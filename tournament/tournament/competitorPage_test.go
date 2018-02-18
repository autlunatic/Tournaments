package tournament

import (
	"reflect"
	"testing"
)

var ttForCompPage = NewTestTournament()

func getWantedCompPageInfo() CompetitorPageInfo {
	out := CompetitorPageInfo{g: ttForCompPage.Groups[1]}

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
		{"TestTournament Bennis Page", args{1, ttForCompPage}, getWantedCompPageInfo()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCompetitorPageInfo(tt.args.competitorID, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToCompetitorPageInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
