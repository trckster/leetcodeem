package main

import "sort"

func canReach(heights []int, index int, bricks int, ladders int) bool {
	increasing := make([]int, 0)

	for i := 0; i < index; i++ {
		if heights[i] < heights[i+1] {
			increasing = append(increasing, heights[i+1]-heights[i])
		}
	}

	sort.Slice(increasing, func(i, j int) bool {
		return increasing[i] < increasing[j]
	})

	heightsSum := 0
	for i := 0; i < len(increasing)-ladders; i++ {
		heightsSum += increasing[i]
	}

	return heightsSum <= bricks
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	l := 0
	r := len(heights)

	for l < r-1 {
		println(l, r)
		m := (l + r) / 2

		if canReach(heights, m, bricks, ladders) {
			l = m
		} else {
			r = m
		}
	}

	return l
}
