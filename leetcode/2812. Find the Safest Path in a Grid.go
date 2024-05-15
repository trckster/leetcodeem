package main

var manhatten [][]int
var visited [][]bool
var currentTrySafety int

func dfs(i, j int) bool {
	visited[i][j] = true

	if manhatten[i][j] < currentTrySafety {
		return false
	}

	if i+1 == len(manhatten) && j+1 == len(manhatten) {
		return true
	}

	result := false

	if i > 0 && !visited[i-1][j] {
		result = result || dfs(i-1, j)
	}

	if j > 0 && !visited[i][j-1] {
		result = result || dfs(i, j-1)
	}

	if i+1 < len(manhatten) && !visited[i+1][j] {
		result = result || dfs(i+1, j)
	}

	if j+1 < len(manhatten) && !visited[i][j+1] {
		result = result || dfs(i, j+1)
	}

	return result
}

func try(k int) bool {
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited); j++ {
			visited[i][j] = false
		}
	}

	currentTrySafety = k

	return dfs(0, 0)
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	manhatten = make([][]int, n)
	visited = make([][]bool, n)
	for i := range grid {
		manhatten[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := range manhatten[i] {
			manhatten[i][j] = -1
		}
	}

	queue := make(map[int]map[int]bool)
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				if _, ok := queue[i]; !ok {
					queue[i] = make(map[int]bool)
				}
				queue[i][j] = true
			}
		}
	}

	currentManhatten := 0
	for len(queue) > 0 {
		newQueue := make(map[int]map[int]bool)
		for i := range queue {
			for j := range queue[i] {
				if manhatten[i][j] == -1 {
					manhatten[i][j] = currentManhatten
				}

				if i > 0 && manhatten[i-1][j] == -1 {
					if _, ok := newQueue[i-1]; !ok {
						newQueue[i-1] = make(map[int]bool)
					}
					newQueue[i-1][j] = true
				}
				if i+1 < n && manhatten[i+1][j] == -1 {
					if _, ok := newQueue[i+1]; !ok {
						newQueue[i+1] = make(map[int]bool)
					}
					newQueue[i+1][j] = true
				}
				if j > 0 && manhatten[i][j-1] == -1 {
					if _, ok := newQueue[i]; !ok {
						newQueue[i] = make(map[int]bool)
					}
					newQueue[i][j-1] = true
				}
				if j+1 < n && manhatten[i][j+1] == -1 {
					if _, ok := newQueue[i]; !ok {
						newQueue[i] = make(map[int]bool)
					}
					newQueue[i][j+1] = true
				}
			}
		}
		queue = newQueue
		currentManhatten++
	}

	l, r := 0, n*2
	for l < r-1 {
		m := (l + r) / 2

		if try(m) {
			l = m
		} else {
			r = m
		}
	}

	return l
}
