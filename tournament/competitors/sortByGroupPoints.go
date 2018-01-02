package competitors

import "reflect"

type sortByGroupPoints struct {
	items []C
}

func (s sortByGroupPoints) Len() int {
	return len(s.items)
}

func (s sortByGroupPoints) Less(i, j int) bool {
	return s.items[i].GetPoints() > s.items[j].GetPoints()
}

func (s sortByGroupPoints) Swap(i, j int) {
	reflect.Swapper(s.items)(i, j)
}
