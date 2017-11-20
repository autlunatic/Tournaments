package tournament

type pairing struct {
	contributor1 Contributor
	contributor2 Contributor
}

func calcPairingsGroup(contributors Contributors) []pairing {
	var result []pairing

	// make a copy of the contributors
	tempContributors := make([]Contributor, len(contributors.items))
	copy(tempContributors, contributors.items)

	// if the count is odd add one dummy contributor for the roundRobin
	if len(contributors.items)%2 > 0 {
		tempContributors = append(tempContributors, Contributor{""})
	}
	// cut the first fixated contributor see roundrobin
	if len(tempContributors) > 1 {
		tempContributors = append(tempContributors[1:])
	}

	// shift one time so it starts with 1v2
	tempContributors = append(tempContributors[1:],tempContributors[0])

	for i := 0; i < len(tempContributors); i++ {
		c1 := contributors.items[0]

		for j := 0; j < (len(tempContributors) / 2)+1; j++ {
			if j == 0 {
				addPair(&result, c1, tempContributors[len(tempContributors)-1])
			} else {
				addPair(&result, tempContributors[j-1], tempContributors[len(tempContributors)-1-j])
			}
		}
		// shift
		tempContributors = append(tempContributors[1:],tempContributors[0])
	}

	/*	for _, c1 := range contributors.items {
			for _, c2 := range tempContributors {
				if c1 != c2 {
					addPair(&result, c1, c2)
				}
			}
		}
	*/
	return result

}
func addPair(pairings *[]pairing, contributor1 Contributor, contributor2 Contributor) {
	if contributor1.name == "" || contributor2.name == "" {
		return
	}
	pair := pairing{contributor1, contributor2}
	*pairings = append(*pairings, pair)
}
