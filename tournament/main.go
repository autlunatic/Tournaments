package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
	"github.com/julienschmidt/httprouter"

	"github.com/autlunatic/Tournaments/tournament/pairings"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/mainpage"
	"github.com/autlunatic/Tournaments/tournament/tournament"
)

var t tournament.T

func main() {
	mux := httprouter.New()

	http.Handle("/", mux)
	mux.GET("/", defaultHandler)
	mux.GET("/gameplan", gamePlanHandler)
	mux.GET("/groups", groupsHandler)
	mux.GET("/results", resultsHandler)
	mux.GET("/mainPage.html", mainPage)
	mux.GET("/inputCompetitors", inputCompetitorsHandler)
	mux.POST("/inputCompetitors", inputCompetitorsHandler)
	mux.GET("/adminPage", adminPageHandler)
	mux.POST("/adminPage", adminPageHandler)
	mux.GET("/inputResults/:id", inputResultHandler)
	mux.POST("/inputResults/:id", inputResultHandler)

	mux.GET("/competitor/:name", competitorPageHandler)

	mux.ServeFiles("/css/*filepath", http.Dir("./css/"))
	mux.Handler("GET", "/favicon.ico", http.NotFoundHandler())
	t.Details = detail.D{
		MinutesAvailForGroupsPhase: 90,
		MinutesPerGame:             15,
		NumberOfParallelGames:      4,
		FinalistCount:              8,
		TournamentStartTime:        time.Now(),
	}
	t.Competitors = competitors.NewTestCompetitors(9)
	t.PairingResults = make(map[int]*pairings.Result)
	t.PointCalcer = tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0)

	t.Build()
	http.ListenAndServe(":8080", nil)
}

func defaultHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	html := "<h1> mainPage </h1>"
	writeHeaderAndHTML(w, html)
}
func mainPage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "mainpage/mainPage.html")
}
func resultsHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	html := pairings.ResultsToHTML(t.Competitors, t.Pairings, t.PairingResults, t.PointCalcer)
	writeHeaderAndHTML(w, html)
}
func groupsHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	html := groups.ToHTML(t.Groups)
	writeHeaderAndHTML(w, html)
}
func gamePlanHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	html := pairings.ToHTML(pairings.CalcedPlanToGamePlan(t.Competitors, t.Plan))
	writeHeaderAndHTML(w, html)
}
func writeHeaderAndHTML(w http.ResponseWriter, html string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, mainpage.ToHTML(html))

}

func adminPageHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var errHTML string
	if req.Method == http.MethodPost {
		minsPerGame, err := strconv.Atoi(req.FormValue("MinutesPerGame"))
		if err != nil {
			errHTML = "Invalid minutes per game value"
		}
		minsTotal, err2 := strconv.Atoi(req.FormValue("MinutesAvailForGroupsPhase"))
		if err2 != nil {
			errHTML = "Invalid total minutes value"
		}
		numberParallel, err3 := strconv.Atoi(req.FormValue("NumberOfParallelGames"))
		if err3 != nil {
			errHTML = "Invalid number of fields"
		}
		finalistCount, err4 := strconv.Atoi(req.FormValue("FinalistCount"))
		if err4 != nil {
		}
		if req.PostFormValue("OK") != "" {
			if errHTML == "" {
				t.Details.MinutesAvailForGroupsPhase = minsTotal
				t.Details.MinutesPerGame = minsPerGame
				t.Details.NumberOfParallelGames = numberParallel
				t.Details.FinalistCount = finalistCount
			}
		} else if req.PostFormValue("build") != "" {
			t.Build()
		} else if req.PostFormValue("calcFinals") != "" {

			t.FinalPairings, err = groups.CalcPairingsForFinals(t.Groups, t.Details.FinalistCount)
			if err != nil {
				t.FinalPairings = []pairings.P{}
			}
			fmt.Println(t.FinalPairings)
		}
	}
	html := tournament.RenderAdminPage(t, errHTML)
	writeHeaderAndHTML(w, html)
}

func inputCompetitorsHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var errHTML string
	if req.Method == http.MethodPost {
		inputTeamName := req.FormValue("competitorName")
		if len(inputTeamName) > 0 {
			err := tryToAddCompetitor(inputTeamName)
			if err != nil {
				errHTML = err.Error()
			} else {
				t.Build()
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
func defaultCSS(w http.ResponseWriter, req *http.Request, _ps httprouter.Params) {
	http.ServeFile(w, req, "default.css")
}

func competitorPageHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	id := competitors.NameToID(t.Competitors, ps.ByName("name"))
	html := tournament.CompetitorPageHTML(id, t)
	writeHeaderAndHTML(w, html)
}

func inputResultHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		id = 0
	}

	p, err2 := t.GetPairingByID(id)
	if err2 != nil {
		writeHeaderAndHTML(w, "<h1> "+err2.Error()+"</h1>")
		return
	}

	var errHTML string
	if req.Method == http.MethodPost {
		ptsC1, err := strconv.Atoi(req.FormValue("competitor1Points"))
		if err != nil {
			errHTML = "Invalid input for Points 1"
		}
		ptsC2, err2 := strconv.Atoi(req.FormValue("competitor2Points"))
		if err2 != nil {
			errHTML = "Invalid input for Points 2"
		}
		if pr, ok := t.PairingResults[p.ID]; ok {
			pr.SetPoints(ptsC1, ptsC2)
		}
		req.Method = http.MethodGet
		competitors.ClearPoints(t.Competitors)
		err = pairings.AddPointsForResults(t.Competitors, t.Pairings, t.PairingResults, t.PointCalcer)
		if err != nil {
			log.Println(err)
		}
		resultsHandler(w, req, ps)

		return
	}
	//t.PairingResults[p.ID].gamePoints1 = 1

	if errHTML != "" {
		writeHeaderAndHTML(w, "<h1> "+errHTML+"</h1>")
		return
	}
	var sif pairings.SimpleInputFormGetter

	html := sif.GetInputForm(t.Competitors, p, t.PairingResults, errHTML)
	writeHeaderAndHTML(w, html)
}
