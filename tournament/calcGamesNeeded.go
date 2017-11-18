package tournament

func calcGroupGamesNeeded(numberOfContributors int) int {
	if numberOfContributors < 2 {
		return 0
	} else {
		return numberOfContributors
	}
}

/*func calcGamesNeeded(numberOfTeams int, numberOfGroups int, numberOfFinalists int) {

}*/
