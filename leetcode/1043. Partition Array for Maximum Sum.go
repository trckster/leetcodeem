package main

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr))

	for i := range dp {
		maxOnNewSubArray := arr[i]
		for j := i; j >= max(i-k+1, 0); j-- {
			maxOnNewSubArray = max(maxOnNewSubArray, arr[j])
			newSubArrayWeight := maxOnNewSubArray * (i - j + 1)
			if j > 0 {
				newSubArrayWeight += dp[j-1]
			}
			dp[i] = max(dp[i], newSubArrayWeight)
		}
	}

	return dp[len(dp)-1]
}
