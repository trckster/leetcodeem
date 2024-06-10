package main

import "sort"

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)

	sort.Ints(expected)

	answer := 0
	for i, h := range heights {
		if expected[i] != h {
			answer++
		}
	}

	return answer
}
