package main

import "slices"

func countSubarrays(nums []int, k int) int64 {
	maximum := slices.Max(nums)
	l, r := 0, 0
	maximumCurrentCount := 0
	for ; l < len(nums) && nums[l] != maximum; l++ {
	}
	for ; r < len(nums); r++ {
		if nums[r] == maximum {
			maximumCurrentCount++
		}

		if maximumCurrentCount == k {
			break
		}
	}

	if maximumCurrentCount < k {
		return 0
	}

	answer := int64(l + 1)
	r++
	for ; r < len(nums); r++ {
		if nums[r] == maximum {
			l++
			for ; l < r && nums[l] != maximum; l++ {

			}
		}

		answer += int64(l + 1)
	}

	return answer
}
