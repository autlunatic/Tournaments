package detail

// D holds Information about the tournament, these are used for serveral calculations and can be seen as a config to the tournament
type D struct {
	NumberOfParallelGames      int
	MinutesPerGame             int
	MinutesAvailForGroupsPhase int
	FinalistCount              int
}

// New creates and initializes Details
func New(numberOfFields int, minutesPerGame int, minAvail int, finalistCount int) *D {
	return &D{MinutesAvailForGroupsPhase: minAvail,
		MinutesPerGame:        minutesPerGame,
		NumberOfParallelGames: numberOfFields}
}
