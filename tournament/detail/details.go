package detail

// D holds Information about the tournament, these are used for serveral calculations and can be seen as a config to the tournament
type D struct {
	NumberOfParallelGames      int
	MinutesPerGame             int
	MinutesAvailForGroupsPhase int
}

// New creates and initializes Details
func New(numberOfFields int, minutesPerGame int, minAvail int) *D {
	d := new(D)
	d.NumberOfParallelGames = numberOfFields
	d.MinutesPerGame = minutesPerGame
	d.MinutesAvailForGroupsPhase = minAvail
	return d
}
