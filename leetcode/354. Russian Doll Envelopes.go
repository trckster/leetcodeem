package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}

		return envelopes[i][0] < envelopes[j][0]
	})

	dp := []int{envelopes[0][1]}

	for i := 1; i < len(envelopes); i++ {
		if envelopes[i][1] > dp[len(dp)-1] {
			dp = append(dp, envelopes[i][1])
			continue
		} else if dp[len(dp)-1] < envelopes[i][1] && dp[len(dp)-1] != envelopes[i][0] {
			dp = append(dp, envelopes[i][1])
		} else {
			l := 0
			r := len(dp) - 1

			for l < r {
				m := (l + r) / 2

				if dp[m] < envelopes[i][1] {
					l = m + 1
				} else {
					r = m
				}
			}

			dp[l] = envelopes[i][1]
		}
	}

	return len(dp)
}

func main() {
	envelopes := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}

	println(maxEnvelopes(envelopes))
}
