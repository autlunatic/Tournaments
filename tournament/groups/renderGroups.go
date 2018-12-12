package groups

import (
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

// CompetitorInfos represents one row in the display of the group
type CompetitorInfos struct {
	Name                string
	GamePoints          int
	GamePointsNegative  int
	GroupPoints         int
	GroupPointsNegative int
}

// GroupInfo ist for representing a group in HTMl
type GroupInfo struct {
	ID    int
	CInfo []CompetitorInfos
}

// GToGroupInfo calculates the GroupInfo for the HTML
func GToGroupInfo(g G) GroupInfo {
	var out GroupInfo
	out.ID = g.ID
	out.CInfo = make([]CompetitorInfos, len(g.Competitors))
	sortedC := competitors.SortedByPoints(g.Competitors)

	for i, ci := range sortedC {
		res, _ := competitors.SumResultPoints(ci.GetResults())
		out.CInfo[i].Name = ci.Name()
		out.CInfo[i].GamePoints = res.GamePoints
		out.CInfo[i].GroupPoints = res.GroupPoints
		out.CInfo[i].GamePointsNegative = res.GamePointsNegative
		out.CInfo[i].GroupPointsNegative = res.GroupPointsNegative
	}
	return out
}

// GetGroupInfos returns all Groupinfos for the G-Slice
func GetGroupInfos(g []G) []GroupInfo {
	var out []GroupInfo
	for _, gi := range g {
		out = append(out, GToGroupInfo(gi))
	}
	return out
}
