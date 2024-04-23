package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int]map[int]bool)
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = make(map[int]bool)
		}
		graph[edge[0]][edge[1]] = true
		if graph[edge[1]] == nil {
			graph[edge[1]] = make(map[int]bool)
		}
		graph[edge[1]][edge[0]] = true
	}

	for len(graph) > 2 {
		leaves := make([]int, 0)
		for k, v := range graph {
			if len(v) == 1 {
				leaves = append(leaves, k)
			}
		}

		for _, leaf := range leaves {
			for k := range graph[leaf] {
				delete(graph[k], leaf)
			}
			delete(graph, leaf)
		}
	}

	answer := make([]int, 0)

	for k := range graph {
		answer = append(answer, k)
	}

	return answer
}
