package pairings

// Result holds data for a Result of a pairing
// the ID in this struct is coupled with the Pairing ID in a tournament plan
type Result struct {
	// gamePoints hold the additional Information about the result
	// in a soccer tournament this would be the goals the tournament goals will be calculated
	GamePoints1 int
	GamePoints2 int
}

// SetPoints Adds
func (r *Result) SetPoints(p1 int, p2 int) {
	r.GamePoints1 = p1
	r.GamePoints2 = p2
}

// Results is a map of Result the key value is the pairing ID
type Results map[int]*Result

// AddResult adds the result to the map and returns the new map
func AddResult(r Results, p1 int, p2 int, pairID int) Results {
	pr, ok := r[pairID]
	if ok {
		pr.SetPoints(p1, p2)
	} else {
		pr = &Result{GamePoints1: p1, GamePoints2: p2}
		r[pairID] = pr
	}
	return r
}
