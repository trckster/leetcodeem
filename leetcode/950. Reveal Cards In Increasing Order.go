package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	answer := make([]int, len(deck))

	skip := false
	for i, j := 0, 0; j < len(deck); i = (i + 1) % len(deck) {
		if answer[i] != 0 {
			continue
		}

		if skip {
			skip = false
			continue
		}

		answer[i] = deck[j]
		j++
		skip = true
	}

	return answer
}
