package tournamentPoints

// TournamentPointCalcer represents an Interface where you can compute the Tournament points out of the game Result
type TournamentPointCalcer interface {
	Calc(gamePoints1 int, gamePoints2 int) (tournamentPoints1 int, TournamentPoints2 int)
}
