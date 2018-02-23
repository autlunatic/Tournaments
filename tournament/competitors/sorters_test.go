package competitors

import (
	"reflect"
	"testing"
)

var testCompetitors []C

func prepareTestCompetitors() {
	testCompetitors = NewTestCompetitors(4)
	ClearPoints(testCompetitors)
	testCompetitors[0].AddPoints(1)
	testCompetitors[0].AddGamePoints(1)
	testCompetitors[0].SetGroupPlacement(2)
	testCompetitors[1].AddPoints(10)
	testCompetitors[1].AddGamePoints(10)
	testCompetitors[1].SetGroupPlacement(1)
	testCompetitors[2].AddPoints(10)
	testCompetitors[2].AddGamePoints(11)
	testCompetitors[2].SetGroupPlacement(2)
	testCompetitors[3].AddPoints(2)
	testCompetitors[3].AddGamePoints(2)
	testCompetitors[3].SetGroupPlacement(1)
}

func wantedSliceByPoints() []C {
	c := make([]C, 4)
	c[0] = testCompetitors[2]
	c[1] = testCompetitors[1]
	c[2] = testCompetitors[3]
	c[3] = testCompetitors[0]
	return c
}

func wantedSlicePlacement() []C {
	c := make([]C, 4)
	c[0] = testCompetitors[1]
	c[1] = testCompetitors[3]
	c[2] = testCompetitors[2]
	c[3] = testCompetitors[0]
	return c
}

func TestGetCompetitorsSortedByGroupPoints(t *testing.T) {
	prepareTestCompetitors()
	tests := []struct {
		name string
		want []C
	}{
		{"sort", wantedSliceByPoints()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedByPoints(testCompetitors); !reflect.DeepEqual(got, tt.want) {
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
			if got := SortedByPlacementAndPoints(testCompetitors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompetitorsSortedByGroupPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
