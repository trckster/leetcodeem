package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)

	answer := int64(0)
	for i := len(happiness) - 1; i >= 0 && k > 0; i, k = i-1, k-1 {
		answer += int64(max(0, happiness[i]-(len(happiness)-i-1)))
	}

	return answer
}
