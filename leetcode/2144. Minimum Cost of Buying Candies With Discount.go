package main

import (
	"slices"
)

func minimumCost(cost []int) int {
	slices.Sort(cost)
	answer := 0

	for i := len(cost) - 1; i >= 0; i-- {
		offset := len(cost) - i

		if offset % 3 == 0 {
			continue
		}

		answer += cost[i]
	}

	return answer
}
