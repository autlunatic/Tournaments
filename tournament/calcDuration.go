package tournament


/*
func main() {
	result := calcDurationMinutes(1, 1, 1)
	fmt.Println("this is a test ", result)
}*/

type TournamentDetails struct {
	numberOfGames int
	numberOfFields int
	minutesPerGame int
}

func NewTournamentDetails(numberOfGames int, numberOfFields int, minutesPerGame int) *TournamentDetails {
	d := new(TournamentDetails)
	d.numberOfFields = numberOfFields
	d.numberOfGames = numberOfGames
	d.minutesPerGame = minutesPerGame
	return d
}

type UseTournamentDetails interface{
	GetTournamentDetails()TournamentDetails
}



// calcs the duration of a tournament in minutes
func InMinutes(tournamentDet UseTournamentDetails) int {
	td := tournamentDet.GetTournamentDetails();
	if td.numberOfFields == 0 {
		return 0
	}

	roundsNeeded := td.numberOfGames / td.numberOfFields
	if td.numberOfGames%td.numberOfFields > 0 {
		roundsNeeded++
	}

	return roundsNeeded * td.minutesPerGame
}
