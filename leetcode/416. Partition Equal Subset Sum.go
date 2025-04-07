package main

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	goal := sum / 2

	dp := make([]bool, goal+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := goal; j >= 1; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}

	return dp[goal]
}
