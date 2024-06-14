package main

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	answer := 0
	currentNumber := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= currentNumber {
			answer += currentNumber - nums[i] + 1
		}

		currentNumber = max(currentNumber+1, nums[i])
	}

	return answer
}
