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
	StartTime     time.Time
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

func CalcedPlanToGamePlan(startTime time.Time, minutesPerGame int, c []competitors.C, cp [][]P) GamePlan {
	var out GamePlan
	for kp := range cp {
		calcedTime := startTime.Add(time.Minute * time.Duration(minutesPerGame*kp))
		for pi := range cp[kp] {
			out.PairingInfo = append(out.PairingInfo,
				PairingInfo{calcedTime,
					calcedTime.Format("15:04"),
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
func ToHTML(gp GamePlan) string {
	tpl := template.Must(template.ParseFiles("pairings/PairingsList.html"))
	var b bytes.Buffer
	tpl.Execute(&b, gp)
	return b.String()
}
