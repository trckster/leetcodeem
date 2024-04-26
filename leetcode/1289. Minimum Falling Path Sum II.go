package main

func minFallingPathSum(grid [][]int) int {
	if len(grid) == 1 {
		return grid[0][0]
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < len(grid[0]); i++ {
		dp[0][i] = grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		minColumn := 0
		for j := 0; j < len(grid[i]); j++ {
			if dp[i-1][minColumn] > dp[i-1][j] {
				minColumn = j
			}
		}
		secondMinColumn := (minColumn + 1) % len(grid[i])
		for j := 0; j < len(grid[i]); j++ {
			if j != minColumn && dp[i-1][secondMinColumn] > dp[i-1][j] {
				secondMinColumn = j
			}
		}

		for j := 0; j < len(grid[i]); j++ {
			if minColumn == j {
				dp[i][j] = grid[i][j] + dp[i-1][secondMinColumn]
			} else {
				dp[i][j] = grid[i][j] + dp[i-1][minColumn]
			}
		}
	}

	answer := dp[len(dp)-1][0]
	for _, n := range dp[len(dp)-1] {
		answer = min(answer, n)
	}

	return answer
}
