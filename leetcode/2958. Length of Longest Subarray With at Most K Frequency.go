package main

func maxSubarrayLength(nums []int, k int) int {
	m := make(map[int]int)

	answer := 0
	for l, r := 0, 0; r < len(nums); r++ {
		m[nums[r]]++

		for m[nums[r]] > k {
			m[nums[l]]--
			l++
		}

		answer = max(answer, r-l+1)
	}

	return answer
}
