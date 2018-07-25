package pairings

import "time"

// AddMinutes adds the given minutes to all pairings
func AddMinutes(p []P, minutes int) {
	for i := range p {
		p[i].StartTime = p[i].StartTime.Add(time.Minute * time.Duration(minutes))
	}
}
