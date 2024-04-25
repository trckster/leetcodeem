package main

func longestIdealString(s string, k int) int {
	if k == 25 {
		return len(s)
	}

	lastOccurrence := make(map[rune]int)
	for i := 'a'; i <= 'z'; i++ {
		lastOccurrence[i] = -1
	}

	dp := make([]int, len(s))
	answer := 0

	for i, c := range s {
		best := -1

		for j := max(c-int32(k), 'a'); j <= min(c+int32(k), 'z'); j++ {
			lastChar := lastOccurrence[j]

			if lastChar == -1 {
				continue
			}

			best = max(best, dp[lastChar])
		}

		if best == -1 {
			dp[i] = 1
		} else {
			dp[i] = best + 1
		}

		answer = max(answer, dp[i])

		lastOccurrence[c] = i
	}

	return answer
}
