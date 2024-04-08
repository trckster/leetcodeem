package main

func countStudents(students []int, sandwiches []int) int {
	m := make(map[int]int)
	for _, s := range students {
		m[s]++
	}

	for _, s := range sandwiches {
		if m[s] == 0 {
			return m[0] + m[1]
		}
		m[s]--
	}

	return 0
}
