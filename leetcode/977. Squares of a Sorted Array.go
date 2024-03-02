package main

import "sort"

func sortedSquares(nums []int) []int {
	for i, n := range nums {
		nums[i] = n * n
	}

	sort.Ints(nums)

	return nums
}
