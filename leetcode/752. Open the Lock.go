package main

import (
	"strconv"
)

func changeDigit(n, shift int, plus bool) int {
	digitToChange := (n / shift) % 10
	n -= digitToChange * shift

	if plus {
		digitToChange = (digitToChange + 1) % 10
	} else {
		digitToChange--
		if digitToChange < 0 {
			digitToChange += 10
		}
	}

	return n + digitToChange*shift
}

func openLock(deadEnds []string, target string) int {
	visited := make(map[int]bool)

	for _, end := range deadEnds {
		number, _ := strconv.Atoi(end)

		if number == 0 {
			return -1
		}

		visited[number] = true
	}

	targetAsInt, _ := strconv.Atoi(target)

	currentStepNumbers := make([]int, 1)
	currentStepNumbers[0] = 0
	step := 0

	for len(currentStepNumbers) > 0 {
		for _, n := range currentStepNumbers {
			if n == targetAsInt {
				return step
			}
		}

		step++
		previousStepNumbersCount := len(currentStepNumbers)

		for i := 0; i < previousStepNumbersCount; i++ {
			previousNumber := currentStepNumbers[i]

			newNumbers := make([]int, 8)
			newNumbers[0] = changeDigit(previousNumber, 1000, true)
			newNumbers[1] = changeDigit(previousNumber, 1000, false)
			newNumbers[2] = changeDigit(previousNumber, 100, true)
			newNumbers[3] = changeDigit(previousNumber, 100, false)
			newNumbers[4] = changeDigit(previousNumber, 10, true)
			newNumbers[5] = changeDigit(previousNumber, 10, false)
			newNumbers[6] = changeDigit(previousNumber, 1, true)
			newNumbers[7] = changeDigit(previousNumber, 1, false)

			for _, n := range newNumbers {
				if !visited[n] {
					currentStepNumbers = append(currentStepNumbers, n)
					visited[n] = true
				}
			}
		}

		currentStepNumbers = currentStepNumbers[previousStepNumbersCount:]
	}

	return -1
}
