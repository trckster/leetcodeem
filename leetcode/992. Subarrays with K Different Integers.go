package main

func subarraysWithKDistinct(nums []int, k int) int {
	l1, l2, differentCount1, differentCount2 := 0, 0, 0, 0
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	answer := 0
	for r := 0; r < len(nums); r++ {
		m1[nums[r]]++
		if m1[nums[r]] == 1 {
			differentCount1++
		}

		m2[nums[r]]++
		if m2[nums[r]] == 1 {
			differentCount2++
		}

		if differentCount1 < k {
			continue
		}

		for differentCount1 > k {
			m1[nums[l1]]--
			if m1[nums[l1]] == 0 {
				differentCount1--
			}
			l1++
		}

		for differentCount2 >= k {
			m2[nums[l2]]--
			if m2[nums[l2]] == 0 {
				differentCount2--
			}
			l2++
		}

		answer += l2 - l1
	}

	return answer
}
