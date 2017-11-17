package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

type testTournamentDet struct{
	numberOfGames int
	numberOfFields int
	minutesPerGame int
}

func (tTD testTournamentDet) GetTournamentDetails() (t TournamentDetails) {
	return TournamentDetails{tTD.numberOfGames,tTD.numberOfFields, tTD.minutesPerGame}
}

func TestAllZero(t *testing.T) {
	result := InMinutes(testTournamentDet{0, 0, 0})
	TestingUtils.CheckEquals(0, result, t)
}

func TestAllOne(t *testing.T) {
	result := InMinutes(testTournamentDet{1, 1, 1})
	TestingUtils.CheckEquals(1, result, t)
}


func TestComplex(t *testing.T) {
	result := InMinutes(testTournamentDet{239,8,15})
	TestingUtils.CheckEquals(450, result, t)

	result = InMinutes(testTournamentDet{10, 2, 15})
	TestingUtils.CheckEquals(75, result, t)

	result = InMinutes(testTournamentDet{6, 1, 5})
	TestingUtils.CheckEquals(30, result, t)

	result = InMinutes(testTournamentDet{9, 8, 15})
	TestingUtils.CheckEquals(30, result, t)

	result = InMinutes(testTournamentDet{10, 3, 1})
	TestingUtils.CheckEquals(4, result, t)
}
