package tournament

import (
	"strconv"
	"testing"
)

func TestCalcRandomDraw(t *testing.T) {

	c := Competitors{}
	c.items = append(c.items,
		newCompetitor("Benni"),
		newCompetitor("Dani"),
		newCompetitor("Zoé"),
		newCompetitor("Mona"))

	calcRandomDraw(c)

	CheckIsAnyDrawNumberDouble(c, t)

}
func CheckIsAnyDrawNumberDouble(c Competitors, t *testing.T) {
	foundDraw := make([]int, len(c.items))
	for _, i := range c.items {
		for _, d := range foundDraw {
			if d == i.drawNumber {
				t.Error("Drawnumbers are not unique -> drawnumber found twice: " + strconv.Itoa(d))
				return
			}
		}
		foundDraw = append(foundDraw, i.drawNumber)

	}
}

func TestIsUniqueInSliceTrue(t *testing.T) {
	slice := []int{1, 2, 3}
	if !isUniqueInSlice(slice, 4) {
		t.Error("4 should be Unique ind slice")
	}

}
func TestIsUniqueInSliceFalse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	if isUniqueInSlice(slice, 4) {
		t.Error("4 should be Unique ind slice")
	}
}
