package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

type calcGroupError struct {
	err string
}

func (e *calcGroupError) Error() string { return e.err }

// CalcGroups adds the competitors to a slice of Group each Group should contain the same amount of competitors if possible
func CalcGroups(c []competitors.C, groupsCount int) (groups []G, err error) {
	if groupsCount*2 > len(c) {
		return groups, &calcGroupError{"too many groups for this count of competitors!"}
	}

	competitorsPerGroup := len(c) / groupsCount
	additionalCompetitors := len(c) % groupsCount

	groups = make([]G, groupsCount)
	for i := range groups {
		groups[i].ID = i + 1
	}

	groupID := 0

	competitorsCountThisGroup := 0
	mc := competitors.SortedByDraw(c)
	for i := 0; i < len(mc); i++ {
		competitorsCountThisGroup++
		groups[groupID].Competitors = append(groups[groupID].Competitors, mc[i])

		if (additionalCompetitors > 0 && competitorsCountThisGroup >= competitorsPerGroup+1) ||
			(additionalCompetitors <= 0 && competitorsCountThisGroup >= competitorsPerGroup) {
			groupID++
			additionalCompetitors--
			competitorsCountThisGroup = 0
		}
	}
	return groups, nil
}
