package tournament

func calcGroups(c CompetitorsGetter, groupsCount int) []group {
	result := make([]group, groupsCount)
	competitors := c.getCompetitors()
	competitorsPerGroup := len(competitors) / groupsCount
	additionalCompetitors := len(competitors) % groupsCount

	for i := range result {
		result[i].Id = i + 1
	}

	groupId := 0

	contributorsCountThisGroup := 0
	for i := 0; i < len(competitors); i++ {
		contributorsCountThisGroup++
		result[groupId].Competitors.items = append(result[groupId].Competitors.items, competitors[i])

		if (additionalCompetitors > 0 && contributorsCountThisGroup >= competitorsPerGroup+1) ||
			(additionalCompetitors <= 0 && contributorsCountThisGroup >= competitorsPerGroup) {
			groupId++
			additionalCompetitors--
			contributorsCountThisGroup = 0
		}
	}
	return result
}
