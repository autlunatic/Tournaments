package groups

import (
	"bytes"
	"html/template"
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
	out.ID = g.id
	out.CInfo = make([]CompetitorInfos, len(g.Competitors))
	for i, ci := range g.Competitors {
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
	gi := GToGroupInfo(g)
	tpl.Execute(&b, gi)
	return b.String()
}
