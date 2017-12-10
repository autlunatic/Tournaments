package pairings

import "github.com/autlunatic/tournaments/tournament/competitors"

type calcGroupPlacements struct {
	pairings      []Pairing
	pairingResult []PairingResult
}

type placement struct {
	competitor competitors.Competitor
	placement  int
}

func calcPlacements() []placement {
	var p []placement

	return p
}
