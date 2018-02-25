package pairings

import (
	"reflect"
	"sort"
)

type sortByIDDesc struct {
	items []P
}

func doSortByIDDesc(ps []P) []P {
	sorter := sortByIDDesc{ps}
	sort.Sort(sorter)
	return sorter.items
}
func (s sortByIDDesc) Len() int {
	return len(s.items)
}

func (s sortByIDDesc) Less(i, j int) bool {
	return s.items[i].ID > s.items[j].ID
}

func (s sortByIDDesc) Swap(i, j int) {
	reflect.Swapper(s.items)(i, j)
}
