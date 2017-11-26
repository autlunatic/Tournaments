package tournament

// calcGroups adds the competitors to a slice of group each group should contain the same amount of competitors if possible
func calcGroups(c CompetitorsGetter, groupsCount int) []group {
	result := make([]group, groupsCount)
	competitors := c.getCompetitors()
	competitorsPerGroup := len(competitors) / groupsCount
	additionalCompetitors := len(competitors) % groupsCount

	for i := range result {
		result[i].id = i + 1
	}

	groupId := 0

	competitorsCountThisGroup := 0
	for i := 0; i < len(competitors); i++ {
		competitorsCountThisGroup++
		result[groupId].competitors.items = append(result[groupId].competitors.items, competitors[i])

		if (additionalCompetitors > 0 && competitorsCountThisGroup >= competitorsPerGroup+1) ||
			(additionalCompetitors <= 0 && competitorsCountThisGroup >= competitorsPerGroup) {
			groupId++
			additionalCompetitors--
			competitorsCountThisGroup = 0
		}
	}
	return result
}
