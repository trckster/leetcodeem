package main

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum := int64(0)
	answer := int64(-1)

	for _, n := range nums {
		if int64(n) < sum {
			answer = max(answer, sum+int64(n))
		}

		sum += int64(n)
	}

	return answer
}
