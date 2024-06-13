package main

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	answer := 0
	for i := 0; i < len(seats); i++ {
		diff := seats[i] - students[i]

		answer += max(diff, -diff)
	}

	return answer
}
