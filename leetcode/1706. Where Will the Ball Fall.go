package main

func emulate(grid [][]int, column int) int {
	x := column

	for y := 0; y < len(grid); y++ {
		direction := grid[y][x]

		if direction == 1 {
			if x == len(grid[0])-1 {
				return -1
			}

			if grid[y][x+1] == -1 {
				return -1
			}

			x++
		} else if direction == -1 {
			if x == 0 {
				return -1
			}

			if grid[y][x-1] == 1 {
				return -1
			}
			x--
		}
	}

	return x
}

func findBall(grid [][]int) []int {
	answer := make([]int, len(grid[0]))

	for i := 0; i < len(grid[0]); i++ {
		answer[i] = emulate(grid, i)
	}

	return answer
}
