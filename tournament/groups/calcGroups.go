package groups

import (
	"github.com/autlunatic/tournaments/tournament/competitors"
)

type calcGroupError struct {
	err string
}

func (e *calcGroupError) Error() string { return e.err }

// CalcGroups adds the competitors to a slice of Group each Group should contain the same amount of competitors if possible
func CalcGroups(c competitors.Getter, groupsCount int) (groups []Group, err error) {
	comps := c.GetCompetitors()
	if groupsCount*2 > len(comps) {
		return groups, &calcGroupError{"too many groups for this count of comps!"}
	}

	competitorsPerGroup := len(comps) / groupsCount
	additionalCompetitors := len(comps) % groupsCount

	groups = make([]Group, groupsCount)
	for i := range groups {
		groups[i].id = i + 1
	}

	groupID := 0

	competitorsCountThisGroup := 0
	for i := 0; i < len(comps); i++ {
		competitorsCountThisGroup++
		groups[groupID].competitors.Items = append(groups[groupID].competitors.Items, comps[i])

		if (additionalCompetitors > 0 && competitorsCountThisGroup >= competitorsPerGroup+1) ||
			(additionalCompetitors <= 0 && competitorsCountThisGroup >= competitorsPerGroup) {
			groupID++
			additionalCompetitors--
			competitorsCountThisGroup = 0
		}
	}
	return groups, nil
}
