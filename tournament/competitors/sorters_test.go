package competitors

import (
	"reflect"
	"testing"
)

var testCompetitors []C

func prepareTestCompetitors() {
	testCompetitors = NewTestCompetitors(4)
	GetCompetitor(testCompetitors, 0).AddPoints(1)
	GetCompetitor(testCompetitors, 0).SetGroupPlacement(2)
	GetCompetitor(testCompetitors, 1).AddPoints(10)
	GetCompetitor(testCompetitors, 1).SetGroupPlacement(1)
	GetCompetitor(testCompetitors, 2).AddPoints(9)
	GetCompetitor(testCompetitors, 2).SetGroupPlacement(2)
	GetCompetitor(testCompetitors, 3).AddPoints(2)
	GetCompetitor(testCompetitors, 3).SetGroupPlacement(1)
}

func wantedSlice() []C {
	c := make([]C, 4)
	c[0] = GetCompetitor(testCompetitors, 1)
	c[1] = GetCompetitor(testCompetitors, 2)
	c[2] = GetCompetitor(testCompetitors, 3)
	c[3] = GetCompetitor(testCompetitors, 0)
	return c
}

func wantedSlicePlacement() []C {
	c := make([]C, 4)
	c[0] = GetCompetitor(testCompetitors, 1)
	c[1] = GetCompetitor(testCompetitors, 3)
	c[2] = GetCompetitor(testCompetitors, 2)
	c[3] = GetCompetitor(testCompetitors, 0)
	return c
}

func TestGetCompetitorsSortedByGroupPoints(t *testing.T) {
	prepareTestCompetitors()
	tests := []struct {
		name string
		want []C
	}{
		{"sort", wantedSlice()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompetitorsSortedByGroupPoints(testCompetitors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompetitorsSortedByGroupPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetCompetitorsSortedByPlacementAndGroupPoints(t *testing.T) {
	prepareTestCompetitors()
	tests := []struct {
		name string
		want []C
	}{
		{"sort", wantedSlicePlacement()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompetitorsSortedByPlacementAndGroupPoints(testCompetitors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompetitorsSortedByGroupPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
