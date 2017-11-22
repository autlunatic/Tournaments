package tournament

import "math/rand"

func calcRandomDraw(c CompetitorsGetter) {
	addedDrawNumbers := make([]int, len(c.getCompetitors()))
	for i := range c.getCompetitors() {
		var d int
		for {
			d = rand.Int()
			if isUniqueInSlice(addedDrawNumbers, d) {
				break
			}
		}
		addedDrawNumbers = append(addedDrawNumbers, d)
		c.getCompetitors()[i].drawNumber = d
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
