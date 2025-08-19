package main

import (
	"math"
)

const c = 1000000007

func intPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func getDPHeight(n, x int) int {
	h := 0
	for i := 0; intPow(i, x) <= n; i++ {
		h++
	}
	return h
}

func buildDP(n, x int) [][]int {
	h := getDPHeight(n, x)

	dp := make([][]int, h)
	dp[0] = make([]int, n+1)
	dp[0][0] = 1

	for i := 1; i < h; i++ {
		dp[i] = make([]int, n+1)
		p := intPow(i, x)

		for j := 0; j <= n; j++ {
			without := dp[i-1][j]
			with := 0
			if p <= j && j-p <= n {
				with = dp[i-1][j-p]
			}

			dp[i][j] = (without + with) % c
		}
	}

	return dp
}

func numberOfWays(n int, x int) int {
	dp := buildDP(n, x)

	return dp[len(dp)-1][n]
}
