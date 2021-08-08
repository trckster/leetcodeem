package main

import "math"

func maxProfit(prices []int) int {
	minimal := math.MaxInt64
	bestResult := 0

	for _, price := range prices {
		if price > minimal && bestResult < price-minimal {
			bestResult = price - minimal
		} else if price < minimal {
			minimal = price
		}
	}

	return bestResult
}

func main() {
	prices := []int{7, 1, 5, 3, 6,4}
	println(maxProfit(prices))
}
