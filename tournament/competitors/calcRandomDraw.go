package competitors

import (
	"math/rand"
	"time"
)

// CalcRandomDraw adds random generated numbers to the DrawNumber of all competitors.
func CalcRandomDraw(c []C) {

	addedDrawNumbers := make([]int, len(c))
	rand.Seed(time.Now().UnixNano())
	for i := range c {
		var d int
		for {
			d = rand.Intn(100000)

			if d < 0 {
				d = -d
			}
			if isUniqueInSlice(addedDrawNumbers, d) {
				break
			}
		}
		addedDrawNumbers = append(addedDrawNumbers, d)
		c[i].SetDrawNumber(d)
	}

}
func isUniqueInSlice(addedDrawNumbers []int, d int) bool {
	for _, a := range addedDrawNumbers {
		if a == d {
			return false
		}
	}
	return true
}
