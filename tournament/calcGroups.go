package tournament


func calcGroups(competitors Competitors, groupsCount int) []group {
	result := make([]group,groupsCount)
	competitorsPerGroup := len(competitors.items) / groupsCount

	gId := 0;
	for _,g := range result{
		g.Id = gId
		gId++
	}

	groupId := -1;

	for i:=0;i<len(competitors.items);i++{
		if i%competitorsPerGroup == 0 {
			groupId++
		}
		result[groupId].Competitors.items = append(result[groupId].Competitors.items, competitors.items[i])
	}

	//result = append(result, group{1,Competitors{ []Competitor{competitors.items[0]}}})
   return result
}