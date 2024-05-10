package main

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractions := make([][]int, len(arr)*(len(arr)-1)/2)

	for i, n := 0, 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fractions[n] = make([]int, 2)
			fractions[n][0] = arr[i]
			fractions[n][1] = arr[j]
			n++
		}
	}

	sort.Slice(fractions, func(i, j int) bool {
		a := float64(fractions[i][0]) / float64(fractions[i][1])
		b := float64(fractions[j][0]) / float64(fractions[j][1])

		return a < b
	})

	answer := make([]int, 2)
	answer[0] = fractions[k-1][0]
	answer[1] = fractions[k-1][1]

	return answer
}
