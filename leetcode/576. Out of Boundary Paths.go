package main

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	board := make([][]int, m)
	board2 := make([][]int, m)
	for i := 0; i < m; i++ {
		board[i] = make([]int, n)
		board2[i] = make([]int, n)
	}
	board[startRow][startColumn] = 1

	currentBoard := board
	nextBoard := board2

	answer := 0
	for move := 0; move < maxMove; move++ {
		if move%2 == 0 {
			currentBoard = board
			nextBoard = board2
		} else {
			currentBoard = board2
			nextBoard = board
		}

		// add to answer
		for i := 0; i < n; i++ {
			answer += currentBoard[0][i]
			answer %= 1_000_000_007
			answer += currentBoard[m-1][i]
			answer %= 1_000_000_007
		}
		for i := 0; i < m; i++ {
			answer += currentBoard[i][0]
			answer %= 1_000_000_007
			answer += currentBoard[i][n-1]
			answer %= 1_000_000_007
		}

		// expansion
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				nextBoard[i][j] = 0
				if i > 0 {
					nextBoard[i][j] += currentBoard[i-1][j]
					nextBoard[i][j] %= 1_000_000_007
				}
				if j > 0 {
					nextBoard[i][j] += currentBoard[i][j-1]
					nextBoard[i][j] %= 1_000_000_007
				}
				if i < m-1 {
					nextBoard[i][j] += currentBoard[i+1][j]
					nextBoard[i][j] %= 1_000_000_007
				}
				if j < n-1 {
					nextBoard[i][j] += currentBoard[i][j+1]
					nextBoard[i][j] %= 1_000_000_007
				}
			}
		}
	}

	return answer
}
