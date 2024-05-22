package main

func isPal(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	partitions := make([][]string, 0)
	currentParts := make([]int, len(s)+1)
	for i := range currentParts {
		currentParts[i] = -1
	}

	var iterateCombinations func(last int)
	iterateCombinations = func(last int) {
		if currentParts[last-1] >= len(s) {
			newAnswer := make([]string, 0)
			for i := 1; i < len(currentParts) && currentParts[i] != -1; i++ {
				from := currentParts[i-1]
				to := currentParts[i]
				newAnswer = append(newAnswer, s[from:to])
			}
			partitions = append(partitions, newAnswer)
		}

		for i := currentParts[last-1]; i < len(s); i++ {
			if isPal(s[currentParts[last-1] : i+1]) {
				currentParts[last] = i + 1
				iterateCombinations(last + 1)
				currentParts[last] = -1
			}
		}
	}

	currentParts[0] = 0
	iterateCombinations(1)

	return partitions
}
