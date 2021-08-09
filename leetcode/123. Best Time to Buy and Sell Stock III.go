package main

import "math"

func maxProfit(prices []int) int {
	rightBest := make([]int, len(prices))

	max := 0
	best := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}

		if max-prices[i] > best {
			best = max - prices[i]
		}

		rightBest[i] = best
	}

	min := math.MaxInt32
	best = rightBest[0]
	leftBest := 0
	for i := 0; i < len(prices) - 1; i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > leftBest {
			leftBest = prices[i] - min
		}

		if leftBest+rightBest[i+1] > best {
			best = leftBest + rightBest[i+1]
		}
	}

	return best
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	println(maxProfit2(prices))

}
