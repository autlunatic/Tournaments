package competitors

import (
	"bytes"
	"html/template"
)

// CompetitorInfo is used for displaying the competitor Name in API oder HTML
type CompetitorInfo struct {
	ID   int
	Name string
}

type inputCompetitors struct {
	C       []C
	ErrHTML string
}

// InputCompetitorsHTML creates the HTML code for the input of the Competitors
func InputCompetitorsHTML(c []C, errHTML string) string {
	tpl := template.Must(template.ParseFiles("competitors/renderCompetitorsInput.html"))
	var b bytes.Buffer
	ic := inputCompetitors{c, errHTML}
	tpl.Execute(&b, ic)
	return b.String()
}

// ToCompetitorInfo converts the competitors for output or API
func ToCompetitorInfo(c []C) []CompetitorInfo {
	var out []CompetitorInfo
	for _, ci := range c {
		out = append(out, CompetitorInfo{ID: ci.ID(),
			Name: ci.Name()})
	}
	return out
}
