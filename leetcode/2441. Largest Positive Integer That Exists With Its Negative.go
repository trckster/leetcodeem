package main

func findMaxK(nums []int) int {
	m := make(map[int]bool)

	answer := -1
	for _, n := range nums {
		if m[-n] {
			answer = max(answer, n, -n)
		}

		m[n] = true
	}

	return answer
}
