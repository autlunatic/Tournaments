package tournament

type pairingResult struct{
	// id refers to the id in the pairingstruct
	pairingID int
	// gamePoints hold the additional Information about the result
	// in a soccer tournament this would be the goals the tournament goals will be calculated
	gamePoints1 int
	gamePoints2 int
}
