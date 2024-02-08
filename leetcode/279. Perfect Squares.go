package main

func numSquares(n int) int {
	squares := make(map[int]bool)
	for i := 1; i*i <= n; i++ {
		squares[i*i] = true
	}

	computed := make([]int, n+1)
	computed[0] = 0

	for i := 1; i <= n; i++ {
		answer := n

		for sq, _ := range squares {
			if i-sq < 0 {
				continue
			}

			answer = min(answer, computed[i-sq]+1)
		}

		computed[i] = answer
	}

	return computed[len(computed)-1]
}
