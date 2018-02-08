package pairings

import "github.com/autlunatic/Tournaments/tournament/competitors"

type simpleInputFormGetter struct{}

func (s simpleInputFormGetter) GetInputForm(c []competitors.C, pairing P, results Results) {

}

// InputFormGetter is the interface that handles the displaying of the page that can input results
type InputFormGetter interface {
	GetInputForm(c []competitors.C, pairing P, results Results) string
}

// InputResultsToHTML shows the input page for results for one pairing ID
func InputResultsToHTML(c []competitors.C, pairing P, results Results, ifg InputFormGetter) string {
	return ifg.GetInputForm(c, pairing, results)
}
