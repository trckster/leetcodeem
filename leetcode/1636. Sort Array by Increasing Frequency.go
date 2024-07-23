package main

import "sort"

func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if m[nums[i]] == m[nums[j]] {
			return nums[i] < nums[j]
		}

		return m[nums[i]] > m[nums[j]]
	})

	return nums
}
