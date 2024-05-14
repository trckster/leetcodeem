package main

func getMaximumGold(grid [][]int) int {
	var dfs func(row, column, gold int) int
	dfs = func(row, column, gold int) int {
		if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[row]) || grid[row][column] <= 0 {
			return gold
		}

		gold += grid[row][column]
		grid[row][column] *= -1

		newGold := max(
			dfs(row+1, column, gold),
			dfs(row, column+1, gold),
			dfs(row, column-1, gold),
			dfs(row-1, column, gold),
		)

		grid[row][column] *= -1

		return newGold
	}

	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			answer = max(answer, dfs(i, j, 0))
		}
	}

	return answer
}
