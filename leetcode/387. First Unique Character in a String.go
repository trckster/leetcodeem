package main

func firstUniqChar(s string) int {
	runeToIndex := make(map[rune]int)

	for i, c := range s {
		if _, ok := runeToIndex[c]; ok {
			runeToIndex[c] = -1
		} else {
			runeToIndex[c] = i
		}
	}

	minIndex := 1_000_000
	for _, i := range runeToIndex {
		if i != -1 {
			minIndex = min(minIndex, i)
		}
	}

	if minIndex == 1_000_000 {
		return -1
	}
	return minIndex
}
