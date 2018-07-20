package tournament

import (
	"bytes"
	"html/template"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/pairings"
)

// CompetitorPageInfo is used in the template for the page of one Competitor
type CompetitorPageInfo struct {
	pairs []pairings.P
	g     groups.G
	ri    []pairings.ResultInfo
}

// CompetitorPageInfoHTML provides the rendered HTML text for the CompetitorPageInfo
type CompetitorPageInfoHTML struct {
	Pairs  template.HTML
	Group  template.HTML
	Result template.HTML
}

func compPageInfoToHTML(cs []competitors.C, cpi CompetitorPageInfo, parallelGames int) CompetitorPageInfoHTML {
	var out CompetitorPageInfoHTML
	gp := pairings.AllPairsToGamePlan(cs, cpi.pairs)
	out.Pairs = template.HTML(pairings.ToHTML("Spielplan", gp))
	out.Group = template.HTML(groups.RenderOneGroup(cpi.g))
	out.Result = template.HTML(pairings.RenderResultInfos(cpi.ri))
	return out
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
func CompetitorPageHTML(competitorID int, t T) string {
	cpi := ToCompetitorPageInfo(competitorID, t)
	cpiH := compPageInfoToHTML(t.Competitors, cpi, t.Details.NumberOfParallelGames)
	tpl := template.Must(template.ParseFiles("tournament/competitorPage.html"))
	var b bytes.Buffer
	tpl.Execute(&b, cpiH)
	return b.String()
}
