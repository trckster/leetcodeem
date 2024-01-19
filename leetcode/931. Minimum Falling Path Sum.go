package main

import "slices"

func minFallingPathSum(matrix [][]int) int {
	for row := 1; row < len(matrix); row++ {
		previousLine := matrix[row-1]
		line := matrix[row]

		for i, value := range line {
			newValue := value + previousLine[i]

			if i+1 < len(matrix) {
				newValue = min(newValue, value+previousLine[i+1])
			}

			if i-1 >= 0 {
				newValue = min(newValue, value+previousLine[i-1])
			}

			line[i] = newValue
		}
	}

	return slices.Min(matrix[len(matrix)-1])
}
