package main

import "sort"

func putMarbles(weights []int, k int) int64 {
	pairedWeights := make([]int, len(weights)-1)
	for i := 0; i < len(weights)-1; i++ {
		pairedWeights[i] = weights[i] + weights[i+1]
	}

	sort.Ints(pairedWeights)

	maxResult := int64(weights[0]) + int64(weights[len(weights)-1])
	minResult := maxResult

	for i := 0; i < k-1; i++ {
		minResult += int64(pairedWeights[i])
	}

	for i := 0; i < k-1; i++ {
		maxResult += int64(pairedWeights[len(pairedWeights)-1-i])
	}

	return maxResult - minResult
}
