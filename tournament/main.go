package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"

	"github.com/autlunatic/Tournaments/tournament/pairings"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
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
	t.Details = detail.D{
		MinutesAvailForGroupsPhase: 90,
		MinutesPerGame:             15,
		NumberOfParallelGames:      4,
	}
	t.Competitors = competitors.NewTestCompetitors(9)
	t.Plan, t.Groups, t.Pairings = groups.CalcMostGamesPerCompetitorPlan(t.Competitors, t.Details)
	http.ListenAndServe(":8080", nil)
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling default")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, mainpage.ToHTML("mainpage"))
}
func mainPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("serving mainpage...")
	http.ServeFile(w, req, "mainpage/mainPage.html")
}
func resultsHandler(w http.ResponseWriter, req *http.Request) {
	html := pairings.ResultsToHTML(t.Competitors, t.Pairings, t.PairingResults, tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, mainpage.ToHTML(html))
}
func groupsHandler(w http.ResponseWriter, req *http.Request) {
	html := groups.ToHTML(t.Groups)
	fmt.Println()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, mainpage.ToHTML(html))
}
func gamePlanHandler(w http.ResponseWriter, req *http.Request) {
	html := pairings.ToHTML(pairings.CalcedPlanToGamePlan(time.Now(), t.Details.MinutesPerGame, t.Competitors, t.Plan))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, mainpage.ToHTML(html))
}

func inputCompetitorsHandler(w http.ResponseWriter, req *http.Request) {
	html := competitors.InputCompetitorsHTML(t.Competitors)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, mainpage.ToHTML(html))
}

func defaultCSS(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "default.css")
}
