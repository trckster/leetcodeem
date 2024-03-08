package main

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int)

	highestFrequency := 0
	for _, n := range nums {
		m[n]++
		highestFrequency = max(highestFrequency, m[n])
	}

	answer := 0
	for _, v := range m {
		if v == highestFrequency {
			answer += v
		}
	}

	return answer
}
