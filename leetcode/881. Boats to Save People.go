package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	l := 0
	r := len(people) - 1

	answer := 0
	for l < r {
		if people[l]+people[r] <= limit {
			l++
		}

		r--
		answer++
	}

	if l == r {
		answer++
	}

	return answer
}
