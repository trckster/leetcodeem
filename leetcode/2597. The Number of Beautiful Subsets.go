package main

import "sort"

func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)

	restrictedPairs := make([]int, 0)
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-n == k {
				newPair := 1<<j | 1<<i
				restrictedPairs = append(restrictedPairs, newPair)
			}
		}
	}

	answer := 0
	for i := 1; i < 1<<len(nums); i++ {
		validAnswer := true
		for _, pair := range restrictedPairs {
			if i&pair == pair {
				validAnswer = false
				break
			}
		}

		if validAnswer {
			answer++
		}
	}

	return answer
}
