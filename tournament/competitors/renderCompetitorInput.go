package competitors

import (
	"bytes"
	"fmt"
	"html/template"
)

// InputCompetitorsHTML creates the HTML code for the input of the Competitors
func InputCompetitorsHTML(c []C) string {
	tpl := template.Must(template.ParseFiles("competitors/renderCompetitorsInput.html"))
	var b bytes.Buffer
	tpl.Execute(&b, c)
	fmt.Println(b.String())
	return b.String()
}
