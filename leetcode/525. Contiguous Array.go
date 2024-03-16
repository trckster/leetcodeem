package main

func findMaxLength(nums []int) int {
	m := make(map[int]int)

	answer := 0
	sum := 0
	m[0] = -1
	for i, n := range nums {
		if n == 0 {
			sum--
		} else {
			sum++
		}

		if v, ok := m[sum]; ok {
			answer = max(answer, i-v)
		} else {
			m[sum] = i
		}
	}

	return answer
}
