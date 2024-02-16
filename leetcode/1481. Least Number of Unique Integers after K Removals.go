package main

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)

	for _, n := range arr {
		if _, ok := m[n]; ok {
			m[n] += 1
		} else {
			m[n] = 1
		}
	}

	counts := make([]int, len(m))

	i := 0
	for _, v := range m {
		counts[i] = v
		i += 1
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] < counts[j]
	})

	answer := len(counts)

	for _, count := range counts {
		k -= count
		if k < 0 {
			break
		}
		answer -= 1
	}

	return answer
}
