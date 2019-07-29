package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/detail"
	"github.com/autlunatic/Tournaments/tournament/groups"
	"github.com/autlunatic/Tournaments/tournament/pairings"
	"github.com/autlunatic/Tournaments/tournament/tournament"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

var t tournament.T
var loadedFromDB = false

func main() {
	fmt.Println("starting tournament server directly (not from cloud engine) ... ")
	apimux := httprouter.New()
	apimux.GET("/api/gamePlan", gamePlanAPI)
	apimux.GET("/api/actualResults", actualResultsAPI)
	apimux.GET("/api/results", resultsAPI)
	apimux.GET("/api/competitors", competitorsAPI)
	apimux.POST("/api/saveResults", resultsInputAPI)
	apimux.POST("/api/saveCompetitors", saveCompetitorsAPI)
	apimux.POST("/api/setDrawNumber", setDrawNumberAPI)
	apimux.POST("/api/saveDetails", saveDetailsAPI)
	apimux.POST("/api/adminFunction", adminFunctionAPI)
	apimux.GET("/api/groups", groupsAPI)
	apimux.GET("/api/getDetails", detailsAPI)
	apimux.GET("/api/isAdmin/:pw", isAdminAPI)
	apimux.GET("/api/isReferee/:pw", isRefereeAPI)
	handler := cors.AllowAll().Handler(apimux)

	t.Details = detail.D{
		MinutesAvailForGroupsPhase: 90,
		MinutesPerGame:             15,
		NumberOfParallelGames:      1,
		FinalistCount:              8,
		TournamentStartTime:        time.Now(),
		FinalsStartTime:            time.Now().Add(time.Minute * 60),
		AdminPassword:              "benni159",
		RefereePassword:            "schiri@4kids",
	}
	t.Competitors = competitors.NewTestCompetitors(9)
	t.PointCalcer = tournamentPoints.NewSimpleTournamentPointCalc(1, 2, 0)

	t.Build()
	http.Handle("/", http.FileServer(http.Dir("./dist/")))
	http.Handle("/adminPage/", http.StripPrefix("/adminPage/", http.FileServer(http.Dir("./dist/"))))
	http.Handle("/gameplan/", http.StripPrefix("/gameplan/", http.FileServer(http.Dir("./dist/"))))
	http.Handle("/results/", http.StripPrefix("/results/", http.FileServer(http.Dir("./dist/"))))
	http.Handle("/mainpage/", http.StripPrefix("/mainpage/", http.FileServer(http.Dir("./dist/"))))
	http.Handle("/groups/", http.StripPrefix("/groups/", http.FileServer(http.Dir("./dist/"))))
	http.HandleFunc("/competitor/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/index.html")
	})
	http.HandleFunc("/inputResults/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/index.html")
	})
	http.Handle("/api/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func fileServerHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println(req)
}

func actualResultsAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	checkLoadFromDB(w, req)
	actual := pairings.FilterActualPairings(t.Competitors, t.Pairings, t.FinalPairings, t.Details)
	old := pairings.FilterOldPairings(t.Competitors, t.Pairings, t.FinalPairings, t.Details)
	coming := pairings.FilterComingPairings(t.Competitors, t.Pairings, t.FinalPairings, t.Details)

	j := pairings.ActualResultsToJSON(t.Competitors, actual, old, coming, t.PairingResults, t.PointCalcer)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, string(j))
}

func checkLoadFromDB(w http.ResponseWriter, req *http.Request) {
	if !loadedFromDB {
		t.LoadFromDataStore(req, w)
		loadedFromDB = true
	}
}
func gamePlanAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	checkLoadFromDB(w, req)
	json := pairings.AllPairsJSON(t.Competitors, t.Pairings, t.FinalPairings, t.Details.NumberOfParallelGames, "Gruppenphase")

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

func isRefereeAPI(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	c := appengine.NewContext(req)
	pw := p.ByName("pw")
	// t.LoadDetails(c)
	log.Infof(c, "referee? %v == %v", pw, t.Details.RefereePassword)

	if pw == t.Details.RefereePassword || t.Details.RefereePassword == "" {
		result, _ := json.Marshal("OK")
		io.WriteString(w, string(result))
		return
	}
	result, _ := json.Marshal("FAIL")
	io.WriteString(w, string(result))
}

func isAdminAPI(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	c := appengine.NewContext(req)
	// t.LoadDetails(c)
	pw := p.ByName("pw")
	log.Infof(c, "admin? %v", pw)

	if pw == t.Details.AdminPassword || t.Details.AdminPassword == "" {
		result, _ := json.Marshal("OK")
		io.WriteString(w, string(result))
		return
	}
	result, _ := json.Marshal("FAIL")
	io.WriteString(w, string(result))
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
	case "calcRandomDraw":
		c := appengine.NewContext(req)
		log.Infof(c, "%v", t.Competitors)
		competitors.CalcRandomDraw(t.Competitors)
		log.Infof(c, "%v", t.Competitors)
	case "addMinutes":
		c := appengine.NewContext(req)
		log.Infof(c, "addmins %v", t.Pairings)
		competitors.CalcRandomDraw(t.Competitors)
		mins, err := strconv.Atoi(function.Params[0])
		if err == nil {
			t.AddMinutesToGroupParingsTime(mins)
		}
		log.Infof(c, "addmins %v mins %v", t.Pairings, mins)
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
	fmt.Println("OK")
}

func setDrawNumberAPI(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var compInfo competitors.CompetitorInfo

	_ = json.NewDecoder(req.Body).Decode(&compInfo)
	for i := range t.Competitors {
		if t.Competitors[i].Name() == compInfo.Name {
			draw, err := strconv.Atoi(compInfo.DrawNumber)
			if err == nil {
				t.Competitors[i].SetDrawNumber(draw)
			}
			return
		}
	}

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
	t.PairingResults = pairings.AddResult(t.PairingResults, ri.Pairing1Pts, ri.Pairing2Pts, p.ID)

	c := appengine.NewContext(req)
	t.SaveResults(c)
	if p.ID < 0 {
		t.RecalcFinals()
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := t.RecalcPoints()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	result, _ := json.Marshal("OK")

	io.WriteString(w, string(result))
	fmt.Println("OK")
}

func calcFin() {
	var err error
	if len(t.FinalPairings) == 0 {
		t.FinalPairings, err = groups.CalcPairingsForFinals(t.Groups, t.Details.FinalistCount, t.Details.NumberOfParallelGames)
		t.SetFinalTimes()
		if err != nil {
			t.FinalPairings = []pairings.P{}
		}
	} else {
		t.RecalcFinals()
	}

}

func tryToAddCompetitor(compName string) error {
	var err error
	t.Competitors, err = competitors.AddByName(t.Competitors, compName)
	return err
}
func defaultCSS(w http.ResponseWriter, req *http.Request, _ps httprouter.Params) {
	http.ServeFile(w, req, "default.css")
}
