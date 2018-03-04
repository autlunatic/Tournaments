package pairings

import (
	"bytes"
	"strconv"
	"text/template"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

// ResultInfo ist the struct for representing Data in the Template
type ResultInfo struct {
	PairingID   int
	PairingInfo string
	Comp1ID     int
	Comp1Name   string
	Comp2ID     int
	Comp2Name   string
	Pairing1Pts int
	Pairing2Pts int
	Group1Pts   int
	Group2Pts   int
	Done        bool
}

// ResultsToResultInfo calculates the struct that is used for presenting the data in HTML
func ResultsToResultInfo(c []competitors.C, p []P, r Results, tpc tournamentPoints.TournamentPointCalcer) []ResultInfo {
	var out []ResultInfo
	for _, pi := range p {
		if res, ok := r[pi.ID]; ok {
			tp1, tp2 := tpc.Calc(res.GamePoints1, res.GamePoints2)
			out = append(out, ResultInfo{
				pi.ID,
				roundToInfo(pi.Round),
				pi.Competitor1ID,
				competitors.GetCompetitor(c, pi.Competitor1ID).Name(),
				pi.Competitor2ID,
				competitors.GetCompetitor(c, pi.Competitor2ID).Name(),
				res.GamePoints1,
				res.GamePoints2,
				tp1,
				tp2,
				true,
			})
		} else {
			out = append(out, ResultInfo{
				pi.ID,
				roundToInfo(pi.Round),
				pi.Competitor1ID,
				competitors.GetCompetitor(c, pi.Competitor1ID).Name(),
				pi.Competitor2ID,
				competitors.GetCompetitor(c, pi.Competitor2ID).Name(),
				0,
				0,
				0,
				0,
				false,
			})

		}
	}
	return out
}

// ResultsToHTML renders the results of a given pairings to HTML
func ResultsToHTML(c []competitors.C, p []P, r Results, tpc tournamentPoints.TournamentPointCalcer) string {
	ri := ResultsToResultInfo(c, p, r, tpc)
	return RenderResultInfos(ri)
}

// RenderResultInfos renders the slice of ResultInfo to HTML
func RenderResultInfos(ri []ResultInfo) string {
	tpl := template.Must(template.ParseFiles("pairings/ResultList.html"))
	var b bytes.Buffer
	tpl.Execute(&b, ri)
	return b.String()
}

// FilterResultInfoByCompID returns resultinfos for one Competitor
func FilterResultInfoByCompID(ris []ResultInfo, compID int) []ResultInfo {
	var out []ResultInfo
	for i, ri := range ris {
		if ri.Comp1ID == compID || ri.Comp2ID == compID {
			out = append(out, ris[i])
		}
	}
	return out
}

// PairingIDToInfo returns the string for the PairingsID
// it formats negative Ids to F[ID] -1 -> F1
func PairingIDToInfo(pID int) string {
	if pID < 0 {
		return "F" + strconv.Itoa(-pID)
	}
	return strconv.Itoa(pID)
}
