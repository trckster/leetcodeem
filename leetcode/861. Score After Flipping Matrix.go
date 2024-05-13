package main

func flipRow(grid [][]int, row int) {
	for i := 0; i < len(grid[row]); i++ {
		grid[row][i] = (grid[row][i] + 1) % 2
	}
}

func flipColumn(grid [][]int, column int) {
	for i := 0; i < len(grid); i++ {
		grid[i][column] = (grid[i][column] + 1) % 2
	}
}

func matrixScore(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			flipRow(grid, i)
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		onesCount := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				onesCount++
			}
		}

		if onesCount < len(grid)-onesCount {
			flipColumn(grid, i)
		}
	}

	answer := 0
	for _, row := range grid {
		n := 0
		for _, bit := range row {
			n <<= 1
			n += bit
		}
		answer += n
	}

	return answer
}
