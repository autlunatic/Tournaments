package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// DetermineFinalists provides functionality to determine who of the groups reached the finals
// count: count of competitors which can reach the finals, semifinals = 4
type determineFinalists struct {
	grps           []G
	count          int
	outIndex       int
	placementIndex int
	out            []competitors.C
}

// DetermineFinalists does the calculation for all groups
func DetermineFinalists(grps []G, count int) []competitors.C {
	d := determineFinalists{
		grps:  grps,
		count: count,
	}

	d.out = make([]competitors.C, d.count)
	for i := 0; i < len(d.grps); i++ {
		d.placementIndex = i
		d.addForPlacement()
	}
	return d.out
}

func (d *determineFinalists) addForPlacement() {
	ssc := d.getSortedCompetitorsForGroup()
	for i := range ssc {
		if d.outIndex >= d.count {
			return
		}
		ssc[i].SetGroupPlacement(d.placementIndex)
		d.out[d.outIndex] = ssc[i]
		d.outIndex++
	}
}

func (d *determineFinalists) getSortedCompetitorsForGroup() (out []competitors.C) {
	var scID int
	sc := make([]competitors.C, len(d.grps))
	for i := range d.grps {
		gc := competitors.GetCompetitorsSortedByGroupPoints(d.grps[i].Competitors)
		sc[scID] = gc[d.placementIndex]
		scID++
	}
	return competitors.GetCompetitorsSortedByGroupPoints(sc)
}
