package pairings

import (
	"bytes"
	"strconv"
	"text/template"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// PairingInfo is used as Struct for the Template that populates one item of the gameplan
type PairingInfo struct {
	FormattedTime string
	Court         string
	Pairing       P
	Comp1Name     string
	Comp2Name     string
}

// GamePlan is used as struct for the Gameplan List.
type GamePlan struct {
	PairingInfo []PairingInfo
}

// CalcedPlanToGamePlan calces the plan for the HTML list
func CalcedPlanToGamePlan(c []competitors.C, cp [][]P) GamePlan {
	var out GamePlan
	for kp := range cp {
		for pi := range cp[kp] {
			out.PairingInfo = append(out.PairingInfo,
				PairingInfo{
					cp[kp][pi].StartTime.Format("15:04"),
					strconv.Itoa(pi + 1),
					cp[kp][pi],
					competitors.GetCompetitor(c, cp[kp][pi].Competitor1ID).Name(),
					competitors.GetCompetitor(c, cp[kp][pi].Competitor2ID).Name(),
				})
		}
	}
	return out
}

// AllPairsToGamePlan returns Gameplan from allpairs for HTML
func AllPairsToGamePlan(c []competitors.C, ap []P) GamePlan {
	var out GamePlan
	for pi := range ap {
		out.PairingInfo = append(out.PairingInfo,
			PairingInfo{
				ap[pi].StartTime.Format("15:04"),
				strconv.Itoa(pi + 1),
				ap[pi],
				competitors.GetCompetitor(c, ap[pi].Competitor1ID).Name(),
				competitors.GetCompetitor(c, ap[pi].Competitor2ID).Name(),
			})
	}
	return out
}

// ToHTML renders the Pairing List to a HTML Page
func ToHTML(description string, gp GamePlan) string {
	htmlData := struct {
		Description string
		Gp          GamePlan
	}{description, gp}
	tpl := template.Must(template.ParseFiles("pairings/PairingsList.html"))
	var b bytes.Buffer
	tpl.Execute(&b, htmlData)
	return b.String()
}
