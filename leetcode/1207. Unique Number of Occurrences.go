package main

func uniqueOccurrences(arr []int) bool {
	counts := make([]int, 2001)
	for _, n := range arr {
		counts[n+1000] += 1
	}

	set := make(map[int]bool)

	for _, count := range counts {
		if count == 0 {
			continue
		}

		if _, ok := set[count]; ok {
			return false
		}

		set[count] = true
	}

	return true
}
