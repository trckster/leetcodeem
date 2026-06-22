func rearrangeCharacters(s string, target string) int {
	needed := make(map[rune]int)
	for _, c := range target {
		needed[c]++
	}

	existing := make(map[rune]int)
	for _, c := range s {
		if _, ok := needed[c]; ok {
			existing[c]++
		}
	}

	for k, v := range needed {
		existing[k] /= v
	}

	min := len(s)
	for _, v := range existing {
		if v < min {
			min = v
		}
	}

	return min
}
