package main

func numIslands(grid [][]byte) int {
	var fillIsland func(i, j int)
	fillIsland = func(i, j int) {
		grid[i][j] = '2'
		if i > 0 && grid[i-1][j] == '1' {
			fillIsland(i-1, j)
		}
		if j > 0 && grid[i][j-1] == '1' {
			fillIsland(i, j-1)
		}
		if i+1 < len(grid) && grid[i+1][j] == '1' {
			fillIsland(i+1, j)
		}
		if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
			fillIsland(i, j+1)
		}
	}

	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				fillIsland(i, j)
				answer++
			}
		}
	}
	return answer
}
