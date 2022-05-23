package main

import "strings"

func findMaxForm(bits []string, maxZeros int, maxOnes int) int {
	dp := make([][]int, maxZeros+1)
	for i := range dp {
		dp[i] = make([]int, maxOnes+1)
	}

	for _, s := range bits {
		ones := strings.Count(s, "1")
		zeros := len(s) - ones

		for i := maxZeros; i >= zeros; i-- {
			for j := maxOnes; j >= ones; j-- {
				applicant := dp[i-zeros][j-ones] + 1

				if applicant > dp[i][j] {
					dp[i][j] = applicant
				}
			}
		}
	}

	return dp[maxZeros][maxOnes]
}

func main() {
	bits := []string{"10", "0001", "111001", "1", "0"}

	println(findMaxForm(bits, 5, 3))
}
