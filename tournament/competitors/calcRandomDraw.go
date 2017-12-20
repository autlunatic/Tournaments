package competitors

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
)

func calcRandomDraw(c Getter) {
	addedDrawNumbers := make([]int64, len(c.GetCompetitors()))
	for i := range c.GetCompetitors() {
		var d int64
		for {
			b := make([]byte, 3)
			_, err := rand.Read(b)
			if err != nil {
				break
			}
			buf := bytes.NewBuffer(b) // b is []byte
			d, err = binary.ReadVarint(buf)
			if err != nil {
				break
			}
			if d < 0 {
				d = -d
			}
			if isUniqueInSlice(addedDrawNumbers, d) {
				break
			}
		}
		addedDrawNumbers = append(addedDrawNumbers, d)
		c.GetCompetitors()[i].SetDrawNumber(d)
	}

}
func isUniqueInSlice(addedDrawNumbers []int64, d int64) bool {
	for _, a := range addedDrawNumbers {
		if a == d {
			return false
		}
	}
	return true
}
