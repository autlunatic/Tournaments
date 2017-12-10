package pairings

// PairingResult holds data for a Result of a pairing
// the ID in this struct is coupled with the Pairing ID in a tournament plan
type PairingResult struct {
	ID int
	// gamePoints hold the additional Information about the result
	// in a soccer tournament this would be the goals the tournament goals will be calculated
	gamePoints1 int
	gamePoints2 int
}
