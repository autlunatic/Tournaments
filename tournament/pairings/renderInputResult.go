package pairings

import (
	"bytes"
	"text/template"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// SimpleInputFormGetter provides an input for 1 Value per Team
type SimpleInputFormGetter struct{}

type simpleInputFields struct {
	ID                int
	Round             int
	Competitor1Name   string
	Competitor2Name   string
	Competitor1Points int
	Competitor2Points int
	ErrHTML           string
}

// GetInputForm implements the InputFormGetter Interface
func (s SimpleInputFormGetter) GetInputForm(c []competitors.C, p P, results Results, errHTML string) string {
	pr, ok := results[p.ID]
	if !ok {
		pr = &Result{0, 0}
		results[p.ID] = pr
	}
	sif := simpleInputFields{p.ID,
		p.Round,
		competitors.GetCompetitor(c, p.Competitor1ID).Name(),
		competitors.GetCompetitor(c, p.Competitor2ID).Name(),
		pr.GamePoints1,
		pr.GamePoints2,
		errHTML,
	}
	tpl := template.Must(template.ParseFiles("pairings/simpleInputForm.html"))
	var b bytes.Buffer
	tpl.Execute(&b, sif)
	return b.String()
}

// InputFormGetter is the interface that handles the displaying of the page that can input results
type InputFormGetter interface {
	GetInputForm(c []competitors.C, pairing P, results Results) string
}

// InputResultsToHTML shows the input page for results for one pairing ID
func InputResultsToHTML(c []competitors.C, pairing P, results Results, ifg InputFormGetter) string {
	return ifg.GetInputForm(c, pairing, results)
}
