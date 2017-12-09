package tournamentPoints


// TournamentPointCalcer represents an Interface where you can compute the Tournament points out of the game Result
type TournamentPointCalcer interface{
	calc(gamePoints1 int, GamePoints2 int) (tournamentPoints1 int, TournamentPoints2 int)
}
