package competitors

import (
	"reflect"
	"sort"
)

// SortedByDraw returns a slice of Competitor which is sorted by the DrawNumber
func SortedByDraw(c []C) []C {
	sorter := &sortByDrawNumber{c}
	sort.Sort(sorter)
	return sorter.items
}

// SortedByPoints returns a slice of Competitor which is sorted by GroupPoints ;)
// High points come first
func SortedByPoints(c []C) []C {
	sorter := &sortByGroupPoints{c}
	sort.Sort(sorter)
	return sorter.items
}

// SortedByPlacementAndPoints returns a slice of Competitor which is sorted first by Placemnt and then by GroupPoints ;)
// Placement 1 comes first, if the placement is the same the one with higher points comes first.
func SortedByPlacementAndPoints(c []C) []C {
	sorter := &sortByPlacementAndGroupPoints{c}
	sort.Sort(sorter)
	return sorter.items
}

type sortByDrawNumber struct {
	items []C
}

func (s sortByDrawNumber) Len() int {
	return len(s.items)
}

func (s sortByDrawNumber) Less(i, j int) bool {
	return s.items[i].DrawNumber() < s.items[j].DrawNumber()
}

func (s sortByDrawNumber) Swap(i, j int) {
	reflect.Swapper(s.items)(i, j)
}

type sortByGroupPoints struct {
	items []C
}

func (s sortByGroupPoints) Len() int {
	return len(s.items)
}

func wonAgainst(is []int, i int) bool {
	for _, w := range is {
		if w == i {
			return true
		}
	}
	return false
}
func lessByGroupPoints(items []C, i, j int) bool {
	resi, woni := SumResultPoints(items[i].GetResults())
	resj, wonj := SumResultPoints(items[j].GetResults())
	if resi.GroupPoints != resj.GroupPoints {
		return resi.GroupPoints > resj.GroupPoints
	}
	if wonAgainst(woni, items[j].ID()) || wonAgainst(wonj, items[i].ID()) {
		return wonAgainst(woni, items[j].ID())
	}
	if resi.GamePoints != resj.GamePoints {
		return resi.GamePoints > resj.GamePoints
	}
	return resi.GamePoints-resi.GamePointsNegative > resj.GamePoints-resj.GamePointsNegative
}

func (s sortByGroupPoints) Less(i, j int) bool {
	return lessByGroupPoints(s.items, i, j)
}

func (s sortByGroupPoints) Swap(i, j int) {
	reflect.Swapper(s.items)(i, j)
}

type sortByPlacementAndGroupPoints struct {
	items []C
}

func (s sortByPlacementAndGroupPoints) Len() int {
	return len(s.items)
}

func (s sortByPlacementAndGroupPoints) Less(i, j int) bool {
	if s.items[i].GroupPlacement() != s.items[j].GroupPlacement() {
		return s.items[i].GroupPlacement() < s.items[j].GroupPlacement()

	}
	return lessByGroupPoints(s.items, i, j)
}

func (s sortByPlacementAndGroupPoints) Swap(i, j int) {
	reflect.Swapper(s.items)(i, j)
}
