package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	l, product := 0, 1

	answer := 0
	for r := 0; r < len(nums); r++ {
		product *= nums[r]

		for l <= r && product >= k {
			product /= nums[l]
			l++
		}

		answer += r - l + 1
	}

	return answer
}
