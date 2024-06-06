package main

import "sort"

type Number struct {
	count int
	index int
}

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	m := make(map[int]*Number)

	for _, n := range nums {
		if number, ok := m[n]; ok {
			number.count++
		} else {
			m[n] = &Number{count: 1}
		}
	}

	newHand := make([]int, len(m))
	i := 0
	for k := range m {
		newHand[i] = k
		i++
	}
	sort.Ints(newHand)
	for k, n := range newHand {
		m[n].index = k
	}

	start := newHand[0]
	nextStart := -1
	lastIndex := -1
	for i = 0; i < len(nums)/k; i++ {
		if nextStart != -1 {
			start = nextStart
			nextStart = -1
		} else {
			if lastIndex+1 >= len(newHand) {
				return false
			}
			start = newHand[lastIndex+1]
		}

		for j := start; j < start+k; j++ {
			number, ok := m[j]
			if !ok || number.count == 0 {
				return false
			}

			number.count--
			lastIndex = number.index

			if number.count > 0 && nextStart == -1 {
				nextStart = j
			}
		}
	}

	return true
}
