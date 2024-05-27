package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		possibleNumber := len(nums) - i

		if nums[i] >= possibleNumber && (i == 0 || nums[i-1] < possibleNumber) {
			return possibleNumber
		}
	}

	return -1
}
