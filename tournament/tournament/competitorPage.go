package tournament

import (
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// CompetitorPageInfo is used in the template for the page of one Competitor
type CompetitorPageInfo struct {
	pairs []pairings.P
	g     groups.G
	ri    []pairings.ResultInfo
}

// ToCompetitorPageInfo calculates the Info for the competitorPage
func ToCompetitorPageInfo(competitorID int, t T) CompetitorPageInfo {
	var out CompetitorPageInfo
	var err error

	out.g, err = groups.GOfCompentitorID(t.Groups, competitorID)
	if err != nil {
		out.g = groups.G{}
	}
	out.pairs = pairings.OfCompetitorID(t.Pairings, competitorID)
	allResults := pairings.ResultsToResultInfo(t.Competitors, t.Pairings, t.PairingResults, t.PointCalcer)
	out.ri = pairings.FilterResultInfoByCompID(allResults, competitorID)

	return out

}

// CompetitorPageHTML returns the page for a Competitor the page includes the following
// next pairings, Group with group points, results of this competitor
func CompetitorPageHTML(t T) {

}
