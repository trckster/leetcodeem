package main

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	countMap := make(map[byte]int)
	for _, v := range tasks {
		countMap[v]++
	}
	count := make([]int, 0)
	for _, v := range countMap {
		count = append(count, v)
	}

	sort.Ints(count)

	fine := (count[len(count)-1] - 1) * n
	othersCount, tail := 0, 0
	for i := len(count) - 2; i >= 0; i-- {
		if count[i] == count[len(count)-1] {
			tail++
		}

		othersCount += count[i]
	}

	return count[len(count)-1] + tail + max(fine, othersCount-tail)
}
