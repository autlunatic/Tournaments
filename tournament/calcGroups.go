package tournament

type calcGroupError struct {
	err string
}

func (e *calcGroupError) Error() string { return e.err }

// calcGroups adds the competitors to a slice of group each group should contain the same amount of competitors if possible
func calcGroups(c CompetitorsGetter, groupsCount int) (groups []group, err error) {
	competitors := c.getCompetitors()
	if groupsCount*2 > len(competitors) {
		return groups, &calcGroupError{"too many groups for this count of competitors!"}
	}

	competitorsPerGroup := len(competitors) / groupsCount
	additionalCompetitors := len(competitors) % groupsCount

	groups = make([]group, groupsCount)
	for i := range groups {
		groups[i].id = i + 1
	}

	groupId := 0

	competitorsCountThisGroup := 0
	for i := 0; i < len(competitors); i++ {
		competitorsCountThisGroup++
		groups[groupId].competitors.items = append(groups[groupId].competitors.items, competitors[i])

		if (additionalCompetitors > 0 && competitorsCountThisGroup >= competitorsPerGroup+1) ||
			(additionalCompetitors <= 0 && competitorsCountThisGroup >= competitorsPerGroup) {
			groupId++
			additionalCompetitors--
			competitorsCountThisGroup = 0
		}
	}
	return groups, nil
}
