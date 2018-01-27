package groups

import (
	"bytes"
	"fmt"
	"html/template"
)

// GroupInfo ist for representing a group in HTMl
type GroupInfo struct {
	ID              int
	CompetitorNames []string
}

// GToGroupInfo calculates the GroupInfo for the HTML
func GToGroupInfo(g []G) []GroupInfo {
	var out []GroupInfo
	for _, gi := range g {

		var grpInf GroupInfo
		grpInf.ID = gi.id
		grpInf.CompetitorNames = make([]string, len(gi.Competitors))
		for i, ci := range gi.Competitors {
			grpInf.CompetitorNames[i] = ci.Name()
		}
		out = append(out, grpInf)
	}
	return out
}

// ToHTML calculates HTML string for representing the groups
func ToHTML(g []G) string {

	tpl := template.Must(template.ParseFiles("groups/renderGroups.html"))
	gi := GToGroupInfo(g)
	var b bytes.Buffer
	tpl.Execute(&b, gi)
	fmt.Println(b.String())
	return b.String()
}
