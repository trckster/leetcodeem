package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make(map[int]map[int]int)
	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make(map[int]int)
		}
		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = make(map[int]int)
		}

		graph[edge[1]][edge[0]] = -1
		graph[edge[0]][edge[1]] = -1
	}

	var setWeight func(root int) int
	setWeight = func(root int) int {
		weight := 1
		for node, _ := range graph[root] {
			delete(graph[node], root)
			graph[root][node] = setWeight(node)
			weight += graph[root][node]
		}
		return weight
	}
	setWeight(0)

	rootAnswer := 0
	var getSumOfDistancesFor func(root, depth int)
	getSumOfDistancesFor = func(root, depth int) {
		for node, _ := range graph[root] {
			rootAnswer += depth
			getSumOfDistancesFor(node, depth+1)
		}
	}
	getSumOfDistancesFor(0, 1)

	answer := make([]int, n)

	var calculateAnswer func(root, currentAnswer int)
	calculateAnswer = func(root, currentAnswer int) {
		answer[root] = currentAnswer
		for node, children := range graph[root] {
			nodeAnswer := currentAnswer - (children - 1) + (n - children - 1)
			calculateAnswer(node, nodeAnswer)
		}
	}
	calculateAnswer(0, rootAnswer)

	return answer
}
