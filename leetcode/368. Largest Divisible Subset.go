package main

import (
	"slices"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	dp := make([]int, len(nums))
	prev := make([]int, len(nums))

	for i, n := range nums {
		prev[i] = -1
		for j := 0; j < i; j++ {
			if n%nums[j] == 0 {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
	}

	maxSize := slices.Max(dp)
	maxSizeIndex := -1

	for i, v := range dp {
		if v == maxSize {
			maxSizeIndex = i
			break
		}
	}

	answer := make([]int, 0)
	for i := maxSizeIndex; i != -1; i = prev[i] {
		answer = append(answer, nums[i])
	}

	return answer
}
