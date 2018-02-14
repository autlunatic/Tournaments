package tournament

import (
	"bytes"
	"html/template"

	"github.com/autlunatic/Tournaments/tournament/detail"
)

// RenderAdminPage returns the page for editing the tournament. Following actions are possible
// rebuild the tournament, calc finals, input Details
func RenderAdminPage(t T, errHTML string) string {
	tplData := struct {
		Det detail.D
		Err string
	}{
		t.Details,
		errHTML,
	}
	tpl := template.Must(template.ParseFiles("tournament/renderTournamentAdministration.html"))
	var b bytes.Buffer
	tpl.Execute(&b, tplData)
	return b.String()
}
