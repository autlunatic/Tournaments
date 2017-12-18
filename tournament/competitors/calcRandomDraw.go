package competitors

import "math/rand"

func calcRandomDraw(c Getter) {
	addedDrawNumbers := make([]int, len(c.GetCompetitors()))
	for i := range c.GetCompetitors() {
		var d int
		for {
			d = rand.Int()
			if isUniqueInSlice(addedDrawNumbers, d) {
				break
			}
		}
		addedDrawNumbers = append(addedDrawNumbers, d)
		c.GetCompetitors()[i].SetDrawNumber(d)
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
