package main

func cherryPickup(grid [][]int) int {
	dp := make([][][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		dp[i] = make([][]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			dp[i][j] = make([]int, len(grid[0]))
			for jj := j; jj < len(grid[0]); jj++ {
				dp[i][j][jj] = -1

				if j > i || jj < len(grid[0])-i-1 {
					continue
				}

				if i == 0 {
					dp[i][j][jj] = grid[i][j] + grid[i][jj]
					continue
				}

				base := dp[i-1][j][jj]

				if j-1 >= 0 {
					base = max(base, dp[i-1][j-1][jj])
				}

				if j+1 < len(grid[0]) {
					base = max(base, dp[i-1][j+1][jj])
				}

				if jj-1 >= 0 {
					base = max(base, dp[i-1][j][jj-1])

					if j-1 >= 0 {
						base = max(base, dp[i-1][j-1][jj-1])
					}

					if j+1 < len(grid[0]) {
						base = max(base, dp[i-1][j+1][jj-1])
					}
				}

				if jj+1 < len(grid[0]) {
					base = max(base, dp[i-1][j][jj+1])

					if j-1 >= 0 {
						base = max(base, dp[i-1][j-1][jj+1])
					}

					if j+1 < len(grid[0]) {
						base = max(base, dp[i-1][j+1][jj+1])
					}
				}

				current := grid[i][j]
				if j != jj {
					current += grid[i][jj]
				}

				dp[i][j][jj] = base + current
			}
		}
	}

	answer := 0
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid[0]); j++ {
			answer = max(answer, dp[len(dp)-1][i][j])
		}
	}
	return answer
}
