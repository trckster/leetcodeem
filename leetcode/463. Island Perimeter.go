package main

func islandPerimeter(grid [][]int) int {
	getPerimeterOfCell := func(i, j int) int {
		p := 0
		if i == 0 || grid[i-1][j] == 0 {
			p++
		}
		if i == len(grid)-1 || grid[i+1][j] == 0 {
			p++
		}
		if j == 0 || grid[i][j-1] == 0 {
			p++
		}
		if j == len(grid[i])-1 || grid[i][j+1] == 0 {
			p++
		}
		return p
	}

	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				answer += getPerimeterOfCell(i, j)
			}
		}
	}
	return answer
}
