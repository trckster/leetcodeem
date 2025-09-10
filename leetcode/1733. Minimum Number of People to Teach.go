package main

import "math"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	knowledgeMap := make(map[int]map[int]bool)
	for i, userLanguages := range languages {
		if knowledgeMap[i] == nil {
			knowledgeMap[i] = make(map[int]bool)
		}

		for _, lang := range userLanguages {
			knowledgeMap[i][lang] = true
		}
	}

	problematic := make(map[int]bool)

	for _, pair := range friendships {
		isProblematic := true

		for _, lang := range languages[pair[0]-1] {
			if knowledgeMap[pair[1]-1][lang] {
				isProblematic = false
				break
			}
		}

		if isProblematic {
			problematic[pair[0]-1] = true
			problematic[pair[1]-1] = true
		}
	}

	bestScore := math.MaxInt32
	for i := 1; i <= n; i++ {
		score := 0

		for user := range problematic {
			if !knowledgeMap[user][i] {
				score++
			}
		}

		if score < bestScore {
			bestScore = score
		}
	}

	return bestScore
}
