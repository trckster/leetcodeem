package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	answer, score, takeI := 0, 0, len(tokens)
	for i := 0; i < takeI; i++ {
		token := tokens[i]

		if token <= power {
			score++
			power -= token
			answer = max(answer, score)
		} else if score >= 1 {
			takeI--
			score--
			power += tokens[takeI]
			i--
		} else {
			return answer
		}
	}

	return answer
}
