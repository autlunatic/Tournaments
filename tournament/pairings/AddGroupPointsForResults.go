package pairings

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

// AddGroupPointsForResults adds the Points to the Competitors groupPoints
func AddGroupPointsForResults(pairings []Pairing, results Results, calcer tournamentPoints.TournamentPointCalcer) error {
	for _, p := range pairings {
		if r, i := results[p.ID]; i {
			p1, p2 := calcer.Calc(r.gamePoints1, r.gamePoints2)
			competitors.GetCompetitor(p.Competitor1ID).AddPoints(p1)
			competitors.GetCompetitor(p.Competitor2ID).AddPoints(p2)
		}
	}
	return nil
}
