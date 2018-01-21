package pairings

import (
	"bytes"
	"strconv"
	"text/template"
	"time"

	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// PairingInfo is used as Struct for the Template that populates one item of the gameplan
type PairingInfo struct {
	startTime time.Time
	court     string
	pairing   P
	Comp1Name string
	Comp2Name string
}

// GamePlan is used as struct for the Gameplan List.
type GamePlan struct {
	pairingInfo []PairingInfo
}

func calcedPlanToGamePlan(startTime time.Time, minutesPerGame int, c []competitors.C, cp [][]P) GamePlan {
	var out GamePlan
	for kp := range cp {
		calcedTime := startTime.Add(time.Minute * time.Duration(minutesPerGame*kp))
		for pi := range cp[kp] {
			out.pairingInfo = append(out.pairingInfo,
				PairingInfo{calcedTime,
					strconv.Itoa(pi + 1),
					cp[kp][pi],
					competitors.GetCompetitor(c, cp[kp][pi].Competitor1ID).Name(),
					competitors.GetCompetitor(c, cp[kp][pi].Competitor2ID).Name(),
				})
		}
	}
	return out
}

// ToHTML renders the Pairing List to a HTML Page
func ToHTML(gp []GamePlan) string {
	tpl := template.Must(template.ParseFiles("PairingsList.html"))
	var b bytes.Buffer
	tpl.Execute(&b, gp)
	return b.String()
}
