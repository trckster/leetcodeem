package main

func exist(board [][]byte, word string) bool {
	var startsFrom func(i, j, k int) bool

	visited := make([]bool, 36)
	startsFrom = func(i, j, k int) bool {
		if k >= len(word) {
			return true
		}

		if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
			return false
		}

		id := j + i*6

		if visited[id] {
			return false
		}

		if board[i][j] != word[k] {
			return false
		}

		visited[id] = true

		result := startsFrom(i+1, j, k+1) || startsFrom(i-1, j, k+1) || startsFrom(i, j+1, k+1) || startsFrom(i, j-1, k+1)

		visited[id] = false

		return result
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if startsFrom(i, j, 0) {
				return true
			}
		}
	}

	return false
}
