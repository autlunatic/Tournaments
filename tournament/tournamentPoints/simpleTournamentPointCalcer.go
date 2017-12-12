package tournamentPoints

// SimpleTournamentPointCalc is an easy to use calculator if the tournament points only depend on win or loss
// winner gets winPoints, loser drawPoints, its not hard to guess what is returned if it is a draw
// this calc implements the TournamentPointCalcer interface
type SimpleTournamentPointCalc struct {
	drawPoints int
	winPoints  int
	losePoints int
}

// NewSimpleTournamentPointCalc creates a SimpleTournamentPointCalc with the given arguments
func NewSimpleTournamentPointCalc(drawPoints int, winPoints int, losePoints int) SimpleTournamentPointCalc {
	return SimpleTournamentPointCalc{drawPoints, winPoints, losePoints}
}

func (s SimpleTournamentPointCalc) Calc(gamePoints1 int, gamePoints2 int) (tournamentPoints1 int, TournamentPoints2 int) {
	if gamePoints1 > gamePoints2 {
		return s.winPoints, s.losePoints
	}
	if gamePoints1 < gamePoints2 {
		return s.losePoints, s.winPoints
	}
	return s.drawPoints, s.drawPoints
}
