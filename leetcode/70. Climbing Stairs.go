package main

func climbStairs(n int) int {
	first := 1
	second := 2

	for ; n > 1; n-- {
		first, second = second, first+second
	}

	return first
}
