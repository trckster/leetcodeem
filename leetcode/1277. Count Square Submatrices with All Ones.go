package main

func countSquares(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	answer := 0

	for i, row := range matrix {
		dp[i] = make([]int, len(row))
		for j, val := range row {
			if val == 0 {
				dp[i][j] = 0
			} else {
				prev := 0
				if i > 0 && j > 0 {
					prev = dp[i-1][j-1]
				}

				top := 0
				if i > 0 {
					top = dp[i-1][j]
				}

				left := 0
				if j > 0 {
					left = dp[i][j-1]
				}

				dp[i][j] = min(prev, top, left) + 1
			}

			answer += dp[i][j]
		}
	}

	return answer
}
