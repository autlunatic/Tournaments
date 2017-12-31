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
	d.addAllBestPlaced(out, &outIndex)
	d.addAllMostPoints(out, &outIndex)
	out = competitors.GetCompetitorsSortedByGroupPoints(out)
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
func (d DetermineFinalists) addAllBestPlaced(out []competitors.Competitor, outIndex *int) {
	for i := range d.grps {
		sc := competitors.GetCompetitorsSortedByGroupPoints(d.grps[i].Competitors)
		out[*outIndex] = sc[0]
		*outIndex++
	}
}

func (d DetermineFinalists) addAllMostPoints(out []competitors.Competitor, outIndex *int) {
	sc := competitors.GetCompetitorsSortedByGroupPoints(d.comps)
	for i := range sc {
		if *outIndex >= d.count {
			return
		}
		if !isInSlice(sc[i], out) {
			out[*outIndex] = sc[i]
			*outIndex++
		}
	}

}
