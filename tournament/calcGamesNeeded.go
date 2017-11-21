package tournament

func calcGroupGamesNeeded(numberOfCompetitors int) int {
	if numberOfCompetitors < 2 {
		return 0
	} else {
		return numberOfCompetitors
	}
}

/*func calcGamesNeeded(numberOfTeams int, numberOfGroups int, numberOfFinalists int) {

}*/
