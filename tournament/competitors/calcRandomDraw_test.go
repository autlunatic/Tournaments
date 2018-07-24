package competitors

import (
	"testing"
)

func TestCalcRandomDrawForNegative(t *testing.T) {
	for i := 0; i < 10; i++ {
		c := make([]C, 0)
		c = append(c,
			New("Benni", 0))

		CalcRandomDraw(c)
		CheckIsNegativeDraw(c, t)
	}
}
func TestCalcRandomDraw(t *testing.T) {

	c := make([]C, 0)
	c = append(c,
		New("Benni", 0),
		New("Dani", 1),
		New("Zoé", 2),
		New("Mona", 3))
	CalcRandomDraw(c)
	CheckIsAnyDrawNumberDouble(c, t)
}

func CheckIsNegativeDraw(c []C, t *testing.T) {
	if c[0].DrawNumber() < 0 {
		t.Errorf("Negative Draw -> %v", c[0].DrawNumber())
	}
}
func CheckIsAnyDrawNumberDouble(c []C, t *testing.T) {
	foundDraw := make([]int, len(c))
	for _, i := range c {
		for _, d := range foundDraw {
			if d == i.DrawNumber() {
				t.Errorf("Drawnumbers are not unique -> drawnumber found twice: %v, comp: %v", d, i)
				return
			}
		}
		foundDraw = append(foundDraw, i.DrawNumber())
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
