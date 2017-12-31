package results

import (
	"reflect"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
)

// DetermineFinalists provides functionality to determine who of the groups reached the finals
// count: count of competitors which can reach the finals, semifinals = 4
type DetermineFinalists struct {
	comps []competitors.Competitor
	grps  []groups.Group
	count int
}

func (d DetermineFinalists) determine() (out []competitors.Competitor) {
	out = make([]competitors.Competitor, d.count)
	var outIndex int
	for i := 0; i < len(d.grps); i++ {
		d.addAllForPlacement(out, &outIndex, i)
	}
	return out
}

func isInSlice(comp competitors.Competitor, comps []competitors.Competitor) bool {
	for i := range comps {
		if reflect.DeepEqual(comp, comps[i]) {
			return true
		}
	}
	return false
}

func (d DetermineFinalists) addAllForPlacement(out []competitors.Competitor, outIndex *int, placementIndex int) {
	sc := make([]competitors.Competitor, len(d.grps))
	ssc := d.getSortedCompetitorsForGroup(sc, placementIndex)
	for i := range ssc {
		if *outIndex >= d.count {
			return
		}
		out[*outIndex] = sc[i]
		*outIndex++
	}
}

func (d DetermineFinalists) getSortedCompetitorsForGroup(sc []competitors.Competitor, placementIndex int) (out []competitors.Competitor) {
	var scID int
	for i := range d.grps {
		gc := competitors.GetCompetitorsSortedByGroupPoints(d.grps[i].Competitors)
		sc[scID] = gc[placementIndex]
		scID++
	}
	ssc := competitors.GetCompetitorsSortedByGroupPoints(sc)
	return ssc
}
