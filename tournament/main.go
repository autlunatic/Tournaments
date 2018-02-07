package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"

	"github.com/autlunatic/Tournaments/tournament/pairings"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/mainpage"
	"github.com/autlunatic/Tournaments/tournament/tournament"
)

var t tournament.T

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/gameplan", gamePlanHandler)
	http.HandleFunc("/groups", groupsHandler)
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("mainPage.html", mainPage)
	http.HandleFunc("/inputCompetitors", inputCompetitorsHandler)
	http.HandleFunc("/default.css", defaultCSS)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	t.Build()
	http.ListenAndServe(":8080", nil)
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	html := "<h1> mainPage </h1>"
	writeHeaderAndHTML(w, html)
}
func mainPage(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "mainpage/mainPage.html")
}
func resultsHandler(w http.ResponseWriter, req *http.Request) {
	html := pairings.ResultsToHTML(t.Competitors, t.Pairings, t.PairingResults, tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0))
	writeHeaderAndHTML(w, html)
}
func groupsHandler(w http.ResponseWriter, req *http.Request) {
	html := groups.ToHTML(t.Groups)
	writeHeaderAndHTML(w, html)
}
func gamePlanHandler(w http.ResponseWriter, req *http.Request) {
	html := pairings.ToHTML(pairings.CalcedPlanToGamePlan(time.Now(), t.Details.MinutesPerGame, t.Competitors, t.Plan))
	writeHeaderAndHTML(w, html)
}
func writeHeaderAndHTML(w http.ResponseWriter, html string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, mainpage.ToHTML(html))

}
func inputCompetitorsHandler(w http.ResponseWriter, req *http.Request) {
	var errHTML string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		inputTeamName := req.FormValue("competitorName")
		if len(inputTeamName) > 0 {
			err := tryToAddCompetitor(inputTeamName)
			if err != nil {
				errHTML = err.Error()
			}
		}
	}
	html := competitors.InputCompetitorsHTML(t.Competitors, errHTML)
	writeHeaderAndHTML(w, html)
}
func tryToAddCompetitor(compName string) error {
	var err error
	t.Competitors, err = competitors.AddByName(t.Competitors, compName)
	return err
}
func defaultCSS(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "default.css")
}
