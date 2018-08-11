package competitors

import (
	"reflect"
	"testing"
)

var testCompetitors []C

func prepareTestCompetitors() {
	testCompetitors = NewTestCompetitors(4)
	ClearPoints(testCompetitors)
	testCompetitors[0].AddResult(ResultPoints{3, 1, 3, 0, 1})
	testCompetitors[0].AddResult(ResultPoints{1, 2, 0, 3, 2})
	testCompetitors[0].AddResult(ResultPoints{1, 2, 0, 3, 3})
	testCompetitors[0].SetGroupPlacement(3)
	testCompetitors[1].AddResult(ResultPoints{1, 3, 0, 3, 0})
	testCompetitors[1].AddResult(ResultPoints{2, 1, 3, 0, 3})
	testCompetitors[1].AddResult(ResultPoints{1, 2, 0, 3, 2})
	testCompetitors[1].SetGroupPlacement(4)
	testCompetitors[2].AddResult(ResultPoints{2, 2, 1, 1, 3})
	testCompetitors[2].AddResult(ResultPoints{2, 1, 3, 0, 0})
	testCompetitors[2].AddResult(ResultPoints{2, 1, 3, 0, 1})
	testCompetitors[2].SetGroupPlacement(2)
	testCompetitors[3].AddResult(ResultPoints{2, 2, 1, 1, 2})
	testCompetitors[3].AddResult(ResultPoints{1, 2, 0, 3, 1})
	testCompetitors[3].AddResult(ResultPoints{2, 1, 3, 0, 0})
	testCompetitors[3].SetGroupPlacement(1)
}

func wantedSliceByPoints() []C {
	c := make([]C, 4)
	c[0] = testCompetitors[2]
	c[1] = testCompetitors[3]
	c[2] = testCompetitors[0]
	c[3] = testCompetitors[1]
	return c
}

func wantedSlicePlacement() []C {
	c := make([]C, 4)
	c[0] = testCompetitors[3]
	c[1] = testCompetitors[2]
	c[2] = testCompetitors[0]
	c[3] = testCompetitors[1]
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
