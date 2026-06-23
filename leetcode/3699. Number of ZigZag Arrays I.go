package main

const MOD = 1000000007

func zigZagArrays(n, l, r int) int {
	size := r - l + 1
	dp := make([]int, size)
	rollingSum := 0
	for i := 0; i < size; i++ {
		dp[i] = i
		rollingSum += i
	}

	newRollingSum := 0
	reverse := false
	for range n - 2 {
		if reverse {
			for i := size - 1; i >= 0; i-- {
				rollingSum = (rollingSum - dp[i] + MOD) % MOD
				dp[i] = rollingSum
				newRollingSum = (newRollingSum + dp[i]) % MOD
			}
		} else {
			for i := 0; i < size; i++ {
				rollingSum = (rollingSum - dp[i] + MOD) % MOD
				dp[i] = rollingSum
				newRollingSum = (newRollingSum + dp[i]) % MOD
			}
		}

		rollingSum = newRollingSum
		newRollingSum = 0
		reverse = !reverse
	}

	answer := 0
	for _, v := range dp {
		answer = (answer + v) % MOD
	}

	return (answer * 2) % MOD
}
