package groups

import (
	"bytes"
	"html/template"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// CompetitorInfos represents one row in the display of the group
type CompetitorInfos struct {
	Name       string
	GamePoints int
	TeamPoints int
}

// GroupInfo ist for representing a group in HTMl
type GroupInfo struct {
	ID    int
	CInfo []CompetitorInfos
}

// GToGroupInfo calculates the GroupInfo for the HTML
func GToGroupInfo(g G) GroupInfo {
	var out GroupInfo
	out.ID = g.ID
	out.CInfo = make([]CompetitorInfos, len(g.Competitors))
	sortedC := competitors.SortedByPlacementAndPoints(g.Competitors)
	for i, ci := range sortedC {
		out.CInfo[i].Name = ci.Name()
		out.CInfo[i].GamePoints = ci.GetGamePoints()
		out.CInfo[i].TeamPoints = ci.GetPoints()
	}
	return out
}

// ToHTML calculates HTML string for representing the groups
func ToHTML(g []G) string {
	tpl := template.Must(template.ParseFiles("groups/renderGroups.html"))
	var groupsStrings []template.HTML
	for _, gi := range g {
		groupsStrings = append(groupsStrings, template.HTML(RenderOneGroup(gi)))
	}
	var b bytes.Buffer
	tpl.Execute(&b, (groupsStrings))
	return b.String()
}

// RenderOneGroup returns HTMl string for one group
func RenderOneGroup(g G) string {
	tpl := template.Must(template.ParseFiles("groups/oneGroup.html"))
	var b bytes.Buffer
	g.Competitors = competitors.SortedByPoints(g.Competitors)
	gi := GToGroupInfo(g)
	tpl.Execute(&b, gi)
	return b.String()
}

// GetGroupInfos returns all Groupinfos for the G-Slice
func GetGroupInfos(g []G) []GroupInfo {
	var out []GroupInfo
	for _, gi := range g {
		out = append(out, GToGroupInfo(gi))
	}
	return out
}
