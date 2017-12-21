package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

type calcGroupError struct {
	err string
}

func (e *calcGroupError) Error() string { return e.err }

// CalcGroups adds the competitors to a slice of Group each Group should contain the same amount of competitors if possible
func CalcGroups(c []competitors.Competitor, groupsCount int) (groups []Group, err error) {
	if groupsCount*2 > len(c) {
		return groups, &calcGroupError{"too many groups for this count of competitors!"}
	}

	competitorsPerGroup := len(c) / groupsCount
	additionalCompetitors := len(c) % groupsCount

	groups = make([]Group, groupsCount)
	for i := range groups {
		groups[i].id = i + 1
	}

	groupID := 0

	competitorsCountThisGroup := 0
	for i := 0; i < len(c); i++ {
		competitorsCountThisGroup++
		groups[groupID].competitors = append(groups[groupID].competitors, c[i])

		if (additionalCompetitors > 0 && competitorsCountThisGroup >= competitorsPerGroup+1) ||
			(additionalCompetitors <= 0 && competitorsCountThisGroup >= competitorsPerGroup) {
			groupID++
			additionalCompetitors--
			competitorsCountThisGroup = 0
		}
	}
	return groups, nil
}
