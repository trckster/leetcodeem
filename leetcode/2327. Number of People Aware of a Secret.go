package main

func peopleAwareOfSecret(n int, delay int, forget int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[delay] = 1

	for i := delay + 1; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-delay]

		if i >= forget {
			dp[i] -= dp[i-forget]
		}

		if dp[i] < 0 {
			dp[i] += 1000000007
		}

		dp[i] %= 1000000007
	}

	answer := 0
	for i := n - 1; i >= n-forget; i-- {
		answer += dp[i]
		answer %= 1000000007
	}

	return answer
}
