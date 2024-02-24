package main

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	dedicated := make(map[int]bool)
	dedicated[0] = true

	meetings = append(meetings, []int{0, firstPerson, 0})
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	for i := 0; i < len(meetings); {
		newGroupMeetingTime := meetings[i][2]
		graph := make(map[int][]int)
		for ; i < len(meetings) && meetings[i][2] == newGroupMeetingTime; i++ {
			m1 := meetings[i][0]
			m2 := meetings[i][1]

			if _, ok := graph[m1]; ok {
				graph[m1] = append(graph[m1], m2)
			} else {
				graph[m1] = make([]int, 1)
				graph[m1][0] = m2
			}
			if _, ok := graph[m2]; ok {
				graph[m2] = append(graph[m2], m1)
			} else {
				graph[m2] = make([]int, 1)
				graph[m2][0] = m1
			}
		}

		isProcessed := make(map[int]bool)
		isProcessedMain := make(map[int]bool)
		var hasDedicated func(int) bool
		var setDedicated func(int)
		hasDedicated = func(node int) bool {
			if _, ok := isProcessed[node]; ok {
				return false
			}
			isProcessedMain[node] = true
			isProcessed[node] = true

			for _, child := range graph[node] {
				if hasDedicated(child) {
					return true
				}
			}

			return dedicated[node]
		}
		setDedicated = func(node int) {
			if _, ok := isProcessed[node]; ok {
				return
			}
			isProcessedMain[node] = true
			isProcessed[node] = true
			dedicated[node] = true
			for _, child := range graph[node] {
				setDedicated(child)
			}
		}
		for node := range graph {
			if _, ok := isProcessedMain[node]; ok {
				continue
			}

			isProcessed = make(map[int]bool)
			if hasDedicated(node) {
				isProcessed = make(map[int]bool)
				setDedicated(node)
			}
		}
	}

	answer := make([]int, len(dedicated))
	i := 0
	for member := range dedicated {
		answer[i] = member
		i++
	}
	return answer
}
