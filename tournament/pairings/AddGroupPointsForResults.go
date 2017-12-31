package pairings

import (
	"fmt"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

// AddGroupPointsForResults adds the Points to the Competitors groupPoints
func AddGroupPointsForResults(c []competitors.Competitor, pairings []Pairing, results Results, calcer tournamentPoints.TournamentPointCalcer) error {
	for _, p := range pairings {
		if r, ok := results[p.ID]; ok {
			p1, p2 := calcer.Calc(r.gamePoints1, r.gamePoints2)
			c1 := competitors.GetCompetitor(c, p.Competitor1ID)
			if c1 != nil {
				c1.AddPoints(p1)
			} else {
				return fmt.Errorf("competitor with ID %v not found", p.Competitor1ID)
			}
			c2 := competitors.GetCompetitor(c, p.Competitor2ID)
			if c2 != nil {
				c2.AddPoints(p2)
			} else {
				return fmt.Errorf("competitor with ID %v not found", p.Competitor2ID)
			}
		}
	}
	return nil
}
