package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/mainpage"
	"github.com/autlunatic/Tournaments/tournament/pairings"
	"github.com/autlunatic/Tournaments/tournament/tournament"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

var t tournament.T

func main() {

	apimux := httprouter.New()
	apimux.GET("/api/gamePlan", gamePlanAPI)
	apimux.GET("/api/results", resultsAPI)
	apimux.GET("/api/competitors", competitorsAPI)
	apimux.POST("/api/saveResults", resultsInputAPI)
	apimux.POST("/api/saveCompetitors", saveCompetitorsAPI)
	apimux.POST("/api/saveDetails", saveDetailsAPI)
	apimux.POST("/api/adminFunction", adminFunctionAPI)
	apimux.GET("/api/groups", groupsAPI)
	apimux.GET("/api/getDetails", detailsAPI)
	handler := cors.AllowAll().Handler(apimux)

	// apimux.GET("/saveResults", resultsInputAPI)
	// go func() {
	// 	http.ListenAndServe(":5050", handler)
	// }()

	// mux := httprouter.New()
	// http.Handle("/", mux)
	// mux.GET("/", defaultHandler)
	// mux.GET("/gameplan", gamePlanHandler)
	// mux.GET("/groups", groupsHandler)
	// mux.GET("/results", resultsHandler)
	// mux.GET("/mainPage.html", mainPage)
	// mux.GET("/inputCompetitors", inputCompetitorsHandler)
	// mux.POST("/inputCompetitors", inputCompetitorsHandler)
	// mux.GET("/adminPage", adminPageHandler)
	// mux.POST("/adminPage", adminPageHandler)
	// mux.GET("/inputResults/:id", inputResultHandler)
	// mux.POST("/inputResults/:id", inputResultHandler)

	// mux.GET("/competitor/:name", competitorPageHandler)

	// mux.ServeFiles("/css/*filepath", http.Dir("./css/"))
	// mux.Handler("GET", "/favicon.ico", http.NotFoundHandler())
	t.Details = detail.D{
		MinutesAvailForGroupsPhase: 90,
		MinutesPerGame:             15,
		NumberOfParallelGames:      1,
		FinalistCount:              8,
		TournamentStartTime:        time.Now(),
		FinalsStartTime:            time.Now().Add(time.Minute * 60),
	}
	t.Competitors = competitors.NewTestCompetitors(9)
	t.PointCalcer = tournamentPoints.NewSimpleTournamentPointCalc(1, 3, 0)

	t.Build()
	// b, err := json.MarshalIndent(t, "  ", "  ")
	// if err == nil {
	// 	 fmt.Println(string(b))
	// } else {
	// 	fmt.Println(err)
	// }

	// mux := http.NewServeMux()
	// mux.Handle("/", http.FileServer(http.Dir("./dist/")))
	// http.ListenAndServe(":8080", mux)
	// fmt.Println("listening ... on :8080")
	http.Handle("/", http.FileServer(http.Dir("./dist/")))
	http.Handle("/adminPage/", http.StripPrefix("/adminPage/", http.FileServer(http.Dir("./dist/"))))
	http.Handle("/api/", handler)
	http.ListenAndServe(":8080", nil)

}

func fileServerHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println(req)

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
	if len(t.FinalPairings) > 0 {
		html += pairings.ResultsToHTML(t.Competitors, t.FinalPairings, t.PairingResults, t.PointCalcer)
	}
	writeHeaderAndHTML(w, html)
}
func groupsHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	html := groups.ToHTML(t.Groups)
	writeHeaderAndHTML(w, html)
}

func gamePlanAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	json := pairings.AllPairsJSON(t.Competitors, t.Pairings, t.FinalPairings, t.Details.NumberOfParallelGames)
	fmt.Println(req)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, json)
}

func competitorsAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	j, _ := json.Marshal(competitors.ToCompetitorInfo(t.Competitors))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, string(j))
}
func resultsAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	j := pairings.ResultsToJSON(t.Competitors, t.Pairings, t.FinalPairings, t.PairingResults, t.PointCalcer)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, string(j))
}
func detailsAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println(t.Details)
	j, _ := json.Marshal(t.Details)
	io.WriteString(w, string(j))
}

func groupsAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	j, _ := json.Marshal(groups.GetGroupInfos(t.Groups))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, string(j))
}

type adminFunction struct {
	Function string
	Params   []string
}

func adminFunctionAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var function adminFunction
	err := json.NewDecoder(req.Body).Decode(&function)
	if err != nil {
		fmt.Println(err)
	}
	// json.NewEncoder(w).Encode(result)
	fmt.Println("request", function)
	var body []byte
	req.Body.Read(body)
	fmt.Println(body)
	fmt.Println(req.Body)
	switch function.Function {
	case "calcFinals":
		calcFin()
		fmt.Println("recalcing")
	case "buildTournament":
		t.Build()
	case "deleteFinalRound":
		t.FinalPairings = []pairings.P{}
	case "saveToDB":
		t.SaveToDataStore(req, w)
	case "loadFromDB":
		t.LoadFromDataStore(req, w)
	}

	result, _ := json.Marshal("OK")
	io.WriteString(w, string(result))
	fmt.Println("OK")
}
func saveDetailsAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var details detail.D
	fmt.Println(t.Details)
	_ = json.NewDecoder(req.Body).Decode(&details)
	// json.NewEncoder(w).Encode(result)
	fmt.Println("request", details)
	t.SetNewDetails(details)

	result, _ := json.Marshal("OK")
	fmt.Println("details", t.Details)
	io.WriteString(w, string(result))
	t.Build()
	fmt.Println("OK")
}

func saveCompetitorsAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var compNames []string
	_ = json.NewDecoder(req.Body).Decode(&compNames)
	// json.NewEncoder(w).Encode(result)
	fmt.Println("request", compNames)
	t.SetNewCompetitors(compNames)

	result, _ := json.Marshal("OK")
	fmt.Println(len(t.Competitors))
	io.WriteString(w, string(result))
	t.Build()
	fmt.Println("OK")
}
func resultsInputAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var ri pairings.ResultInfo
	_ = json.NewDecoder(req.Body).Decode(&ri)
	// json.NewEncoder(w).Encode(result)
	fmt.Println("request", ri)

	p, err2 := t.GetPairingByID(ri.PairingID)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(p)
	fmt.Println(p.ID)
	fmt.Println(len(t.PairingResults))
	pr, ok := t.PairingResults[p.ID]
	if !ok {
		pr = &pairings.Result{GamePoints1: ri.Pairing1Pts, GamePoints2: ri.Pairing2Pts}
		t.PairingResults[p.ID] = pr
	}
	pr.SetPoints(ri.Pairing1Pts, ri.Pairing2Pts)
	if p.ID < 0 {
		t.RecalcFinals()
	}
	fmt.Println("pr", pr)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	competitors.ClearPoints(t.Competitors)
	err := pairings.AddPointsForResults(t.Competitors, t.Pairings, t.PairingResults, t.PointCalcer)
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}
	result, _ := json.Marshal("OK")

	io.WriteString(w, string(result))
	fmt.Println("OK")
}

func gamePlanHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	html := pairings.ToHTML("Spielplan", pairings.AllPairsToGamePlan(t.Competitors, t.Pairings, t.Details.NumberOfParallelGames))
	if len(t.FinalPairings) > 0 {
		html = html + pairings.ToHTML("Finalrunden", pairings.AllPairsToGamePlan(t.Competitors, t.FinalPairings, t.Details.NumberOfParallelGames))
	}
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
			t.SaveToDataStore(req, w)
		} else if req.PostFormValue("calcFinals") != "" {
			calcFin()
		} else if req.PostFormValue("deleteFinals") != "" {
			t.FinalPairings = []pairings.P{}

		}
	}
	html := tournament.RenderAdminPage(t, errHTML)
	writeHeaderAndHTML(w, html)
}
func calcFin() {
	var err error
	if len(t.FinalPairings) == 0 {
		t.FinalPairings, err = groups.CalcPairingsForFinals(t.Groups, t.Details.FinalistCount)
		t.SetFinalTimes()
		if err != nil {
			t.FinalPairings = []pairings.P{}
		}
	} else {
		t.RecalcFinals()
	}

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
	fmt.Println(len(t.PairingResults))
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
		fmt.Println(len(t.PairingResults))
		if pr, ok := t.PairingResults[p.ID]; ok {
			pr.SetPoints(ptsC1, ptsC2)
			if id < 0 {
				t.RecalcFinals()
			}
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

	if errHTML != "" {
		writeHeaderAndHTML(w, "<h1> "+errHTML+"</h1>")
		return
	}
	var sif pairings.SimpleInputFormGetter

	html := sif.GetInputForm(t.Competitors, p, t.PairingResults, errHTML)
	writeHeaderAndHTML(w, html)
}
