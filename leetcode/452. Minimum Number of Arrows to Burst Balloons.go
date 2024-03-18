package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	answer := 1
	lastArrow := points[0][1]
	for _, point := range points {
		if point[0] > lastArrow {
			answer++
			lastArrow = point[1]
		}
	}

	return answer
}
