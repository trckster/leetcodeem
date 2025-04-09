package main

import "sort"

func minOperations(nums []int, k int) int {
	set := make(map[int]any)
	for _, num := range nums {
		set[num] = struct{}{}
	}

	uniqueNumbers := make([]int, 0, len(set))
	for num := range set {
		uniqueNumbers = append(uniqueNumbers, num)
	}

	sort.Ints(uniqueNumbers)

	if k > uniqueNumbers[0] {
		return -1
	} else if k < uniqueNumbers[0] {
		return len(uniqueNumbers)
	}

	return len(uniqueNumbers) - 1
}
