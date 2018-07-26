package pairings

import (
	"fmt"

	"github.com/autlunatic/Tournaments/tournament/competitors"
	"github.com/autlunatic/Tournaments/tournament/tournamentPoints"
)

// AddPointsForResults adds the Points to the Competitors groupPoints
func AddPointsForResults(c []competitors.C, pairings []P, results Results, calcer tournamentPoints.TournamentPointCalcer) error {
	for _, p := range pairings {
		if r, ok := results[p.ID]; ok {
			r2 := *r
			p1, p2 := calcer.Calc(r2.GamePoints1, r2.GamePoints2)
			c1 := competitors.GetCompetitor(c, p.Competitor1ID)
			if c1.ID() > -1 {
				c1.AddResult(competitors.ResultPoints{GroupPoints: p1,
					GroupPointsNegative: p2,
					GamePoints:          r2.GamePoints1,
					GamePointsNegative:  r2.GamePoints2,
					AgainstCompetitorID: p.Competitor2ID})
			} else {
				return fmt.Errorf("competitor with ID %v not found", p.Competitor1ID)
			}
			c2 := competitors.GetCompetitor(c, p.Competitor2ID)
			if c2.ID() > -1 {
				c2.AddResult(competitors.ResultPoints{GroupPoints: p2,
					GroupPointsNegative: p1,
					GamePoints:          r2.GamePoints2,
					GamePointsNegative:  r2.GamePoints1,
					AgainstCompetitorID: p.Competitor1ID})
			} else {
				return fmt.Errorf("competitor with ID %v not found", p.Competitor2ID)
			}
		}
	}
	return nil
}
