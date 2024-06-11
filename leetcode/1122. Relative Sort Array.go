package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	indexes := make(map[int]int)
	for i, n := range arr2 {
		indexes[n] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		indexA := 1001 + arr1[i]
		if v, ok := indexes[arr1[i]]; ok {
			indexA = v
		}

		indexB := 1001 + arr1[j]
		if v, ok := indexes[arr1[j]]; ok {
			indexB = v
		}

		return indexA < indexB
	})

	return arr1
}
