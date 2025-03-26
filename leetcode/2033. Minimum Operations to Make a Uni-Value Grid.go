package main

import (
	"math"
)

func minOperations(grid [][]int, x int) int {
	unitedRemain := grid[0][0] % x

	sum := 0
	minimum := math.MaxInt64
	maximum := math.MinInt64
	for _, row := range grid {
		for _, cell := range row {
			if cell%x != unitedRemain {
				return -1
			}

			sum += cell
			if cell < minimum {
				minimum = cell
			}
			if cell > maximum {
				maximum = cell
			}
		}
	}

	jumpsCount := func(target int) int {
		jumps := 0

		for _, row := range grid {
			for _, cell := range row {
				jumps += int(math.Abs(float64(target)-float64(cell))) / x
			}
		}

		return jumps
	}

	l := 0
	r := (maximum - minimum) / x

	if l == r {
		return 0
	}

	absolute := min(jumpsCount(minimum), jumpsCount(maximum))

	for {
		m1 := l + (r-l)/3
		m2 := r - (r-l)/3
		if (r-l)/3 == 0 {
			m1++
		}

		m1c := jumpsCount(minimum + m1*x)
		m2c := jumpsCount(minimum + m2*x)

		absolute = min(absolute, m1c, m2c)

		if m2-m1 <= 1 {
			return absolute
		}

		if m1c >= m2c {
			l = m1
		} else {
			r = m2
		}
	}
}
