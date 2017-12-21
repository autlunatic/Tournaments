package competitors

import (
	"testing"
)

func TestCalcRandomDraw(t *testing.T) {

	c := make([]Competitor, 0)
	c = append(c,
		NewCompetitor("Benni", 0),
		NewCompetitor("Dani", 1),
		NewCompetitor("Zoé", 2),
		NewCompetitor("Mona", 3))
	CalcRandomDraw(c)
	CheckIsAnyDrawNumberDouble(c, t)
}

func CheckIsAnyDrawNumberDouble(c []Competitor, t *testing.T) {
	foundDraw := make([]int64, len(c))
	for _, i := range c {
		for _, d := range foundDraw {
			if d == i.DrawNumber() {
				t.Errorf("Drawnumbers are not unique -> drawnumber found twice: %v", d)
				return
			}
		}
		foundDraw = append(foundDraw, i.DrawNumber())
	}
}

func TestIsUniqueInSliceTrue(t *testing.T) {
	slice := []int64{1, 2, 3}
	if !isUniqueInSlice(slice, 4) {
		t.Error("4 should be Unique ind slice")
	}
}

func TestIsUniqueInSliceFalse(t *testing.T) {
	slice := []int64{1, 2, 3, 4, 5}
	if isUniqueInSlice(slice, 4) {
		t.Error("4 should be Unique ind slice")
	}
}
