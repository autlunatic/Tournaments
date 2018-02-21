package pairings

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func TestCalcPairings(t *testing.T) {

	benni := competitors.New("Benni", 0)
	dani := competitors.New("Dani", 1)
	cs := make([]competitors.C, 2)
	cs[0] = benni
	cs[1] = dani

	pairings, _ := CalcPairings(cs, 1)

	TestingUtils.CheckEquals(1, len(pairings), "", t)
	pair := pairings[0]

	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
}
func TestCalcPairingsEmptyCompetitorsShouldNotPanic(t *testing.T) {
	var cs []competitors.C
	_, err := CalcPairings(cs, 1)
	if err == nil {
		t.Error("expected Error didnt return!")
	}
}

func TestCalcPairings3Competitors(t *testing.T) {
	benni := competitors.New("Benni", 0)
	dani := competitors.New("Dani", 1)
	zoe := competitors.New("Zoé", 2)
	cs := make([]competitors.C, 3)
	cs[0] = benni
	cs[1] = dani
	cs[2] = zoe

	pairings, _ := CalcPairings(cs, 1)

	TestingUtils.CheckEquals(3, len(pairings), "", t)
	pair := pairings[0]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	pair = pairings[1]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())

	}
	pair = pairings[2]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Dani") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Zoé") {
		t.Errorf("first competitor should be Dani (but was %s) and second should be Zoé (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
}

func checkPairingDoubles(pairings []P) (msg string) {
	var count int
	for _, pair := range pairings {
		count = 0
		for _, pair2 := range pairings {
			if isSamePair(pair, pair2) {
				count++
			}
		}
		if count > 1 {
			return fmt.Sprintf("CompetitorIds: %v vs. %v", pair.Competitor1ID, pair.Competitor2ID)
		}

	}
	return ""
}
func isSamePair(p P, p2 P) bool {
	return (p.Competitor1ID == p2.Competitor1ID) && (p.Competitor2ID == p2.Competitor2ID) ||
		(p.Competitor1ID == p2.Competitor2ID) && (p.Competitor2ID == p2.Competitor1ID)

}

func TestCalcPairings5Competitors(t *testing.T) {
	var cs []competitors.C
	cs = append(cs,
		competitors.New("Benni", 0),
		competitors.New("Dani", 1),
		competitors.New("Zoé", 2),
		competitors.New("Mona", 3),
		competitors.New("Andrea", 4))

	pairings, _ := CalcPairings(cs, 1)

	TestingUtils.CheckEquals(10, len(pairings), "", t)
	pair := pairings[0]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Dani") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Dani (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	pair = pairings[1]
	if (competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Mona") || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Andrea") {
		t.Errorf("first competitor should be Mona (but was %s) and second should be Andrea (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	pair = pairings[2]
	if competitors.GetCompetitor(cs, pair.Competitor1ID).Name() != "Benni" || (competitors.GetCompetitor(cs, pair.Competitor2ID).Name() != "Zoé") {
		t.Errorf("first competitor should be Benni (but was %s) and second should be Zoé (but was %s)",
			competitors.GetCompetitor(cs, pair.Competitor1ID).Name(), competitors.GetCompetitor(cs, pair.Competitor2ID).Name())
	}
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same c found! " + msg)
	}
}
func TestCalcPairings12Competitors(t *testing.T) {
	cs := competitors.NewTestCompetitors(12)

	pairings, _ := CalcPairings(cs, 1)
	TestingUtils.CheckEquals(66, len(pairings), "", t)
	TestingUtils.CheckEquals(6, pairings[33].Round, "Round", t)
	if msg := checkPairingDoubles(pairings); msg != "" {
		t.Errorf("same competitors found! " + msg)
	}
}

func getPairingsForTestMaxRound() []P {
	p := make([]P, 3)
	p[0] = P{1, 2, 3, 1, 1}
	p[1] = P{3, 4, 6, 1, 1}
	p[2] = P{1, 2, 5, 1, 1}
	return p
}
func TestGetMaxRoundOfPairings(t *testing.T) {
	tests := []struct {
		name string
		args []P
		want int
	}{
		{"Wanted Maxround 6", getPairingsForTestMaxRound(), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxRoundOfPairings(tt.args); got != tt.want {
				t.Errorf("GetMaxRoundOfPairings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairing_ToString(t *testing.T) {
	type fields struct {
		Competitor1ID int
		Competitor2ID int
		Round         int
		ID            int
		GroupID       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"simple ToStringTest", fields{1, 2, 3, 1, 1}, "round: 3; 1 vs. 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := P{
				Competitor1ID: tt.fields.Competitor1ID,
				Competitor2ID: tt.fields.Competitor2ID,
				Round:         tt.fields.Round,
				ID:            tt.fields.ID,
				GroupID:       tt.fields.GroupID,
			}
			if got := p.ToString(); got != tt.want {
				t.Errorf("Pairing.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairing_InPairings(t *testing.T) {

	type fields struct {
		Competitor1ID int
		Competitor2ID int
		Round         int
		ID            int
		GroupID       int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"is in slice", fields{1, 2, 1, 1, 1}, true},
		{"is not in slice", fields{7, 8, 1, 4, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := P{
				Competitor1ID: tt.fields.Competitor1ID,
				Competitor2ID: tt.fields.Competitor2ID,
				Round:         tt.fields.Round,
				ID:            tt.fields.ID,
				GroupID:       tt.fields.GroupID,
			}
			ps := make([]P, 3)
			ps[0] = P{1, 2, 1, 1, 1}
			ps[1] = P{3, 4, 1, 2, 1}
			ps[2] = P{5, 6, 1, 3, 1}
			if got := p.InPairings(ps); got != tt.want {
				t.Errorf("Pairing.InPairings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfCompetitorID(t *testing.T) {
	type args struct {
		ps     []P
		compID int
	}
	tests := []struct {
		name string
		args args
		want []P
	}{
		{"simple with 3 pairings",
			args{ps: []P{{1, 2, 1, 1, 1},
				{1, 3, 2, 2, 1},
				{2, 3, 3, 3, 1},
			}, compID: 1},
			[]P{{1, 2, 1, 1, 1},
				{1, 3, 2, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OfCompetitorID(tt.args.ps, tt.args.compID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfCompetitorID() = %v, want %v", got, tt.want)
			}
		})
	}
}
