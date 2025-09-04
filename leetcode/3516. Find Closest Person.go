package main

import "math"

func findClosest(x int, y int, z int) int {
	diffX := math.Abs(float64(x) - float64(z))
	diffY := math.Abs(float64(y) - float64(z))

	if diffX < diffY {
		return 1
	} else if diffX > diffY {
		return 2
	}

	return 0
}
