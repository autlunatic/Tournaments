package tournament

type pairing struct {
	contributor1 *Contributor
	contributor2 *Contributor
}

func calcPairingsGroup(contributors Contributors) []pairing {
	var result []pairing
	for _, c1 := range contributors.items {
		for _, c2 := range contributors.items {
			if c1 != c2 {
				addpair(&result, &c1, &c2)
			}
		}
	}

	return result

}
func addpair(pairings *[]pairing, contributor *Contributor, contributor2 *Contributor) {
	pair := pairing{contributor, contributor2}
	*pairings = append(*pairings, pair)
}
