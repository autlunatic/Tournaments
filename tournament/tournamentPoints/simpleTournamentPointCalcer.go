package tournamentPoints


type SimpleTournamentPointCalc struct{
	pointsForDraw int
	pointsForWin int
	pointsForLose int
}

func (s SimpleTournamentPointCalc) calc(gamePoints1 int, GamePoints2 int) (tournamentPoints1 int, TournamentPoints2 int) {
	if gamePoints1>GamePoints2{
		return s.pointsForWin,s.pointsForLose
	}
	if gamePoints1<GamePoints2{
		return s.pointsForLose,s.pointsForWin
	}
	return s.pointsForDraw,s.pointsForDraw
}
