package main

func largestLocal(grid [][]int) [][]int {
	answer := make([][]int, len(grid)-2)

	for i := 1; i+1 < len(grid); i++ {
		answer[i-1] = make([]int, len(grid)-2)
		for j := 1; j+1 < len(grid); j++ {
			answer[i-1][j-1] = max(
				grid[i-1][j-1],
				grid[i-1][j],
				grid[i-1][j+1],
				grid[i][j-1],
				grid[i][j],
				grid[i][j+1],
				grid[i+1][j-1],
				grid[i+1][j],
				grid[i+1][j+1],
			)
		}
	}

	return answer
}
