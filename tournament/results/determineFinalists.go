package results

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
)

// DetermineFinalists provides functionality to determine who of the groups reached the finals
// count: count of competitors which can reach the finals, semifinals = 4
type DetermineFinalists struct {
	comps          []competitors.Competitor
	grps           []groups.Group
	count          int
	outIndex       int
	placementIndex int
	out            []competitors.Competitor
}

func (d *DetermineFinalists) determine() []competitors.Competitor {
	d.out = make([]competitors.Competitor, d.count)
	for i := 0; i < len(d.grps); i++ {
		d.placementIndex = i
		d.addForPlacement()
	}
	return d.out
}

func (d *DetermineFinalists) addForPlacement() {
	ssc := d.getSortedCompetitorsForGroup()
	for i := range ssc {
		if d.outIndex >= d.count {
			return
		}
		d.out[d.outIndex] = ssc[i]
		d.outIndex++
	}
}

func (d *DetermineFinalists) getSortedCompetitorsForGroup() (out []competitors.Competitor) {
	var scID int
	sc := make([]competitors.Competitor, len(d.grps))
	for i := range d.grps {
		gc := competitors.GetCompetitorsSortedByGroupPoints(d.grps[i].Competitors)
		sc[scID] = gc[d.placementIndex]
		scID++
	}
	return competitors.GetCompetitorsSortedByGroupPoints(sc)
}
