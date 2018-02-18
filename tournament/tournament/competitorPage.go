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
	return CompetitorPageInfo{}

}

// CompetitorPageHTML returns the page for a Competitor the page includes the following
// next pairings, Group with group points, results of this competitor
func CompetitorPageHTML(t T) {

}
