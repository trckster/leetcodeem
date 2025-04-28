package main

func countSubarrays(nums []int, k int64) int64 {
	answer := int64(0)
	currentSum := int64(0)

	for right, left := 0, 0; right < len(nums); right++ {
		currentSum += int64(nums[right])

		for currentSum*int64(right-left+1) >= k {
			currentSum -= int64(nums[left])
			left++
		}

		if currentSum > 0 {
			answer += int64(right - left + 1)
		}
	}

	return answer
}
