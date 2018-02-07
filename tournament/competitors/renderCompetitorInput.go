package competitors

import (
	"bytes"
	"html/template"
)

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
