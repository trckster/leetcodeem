package main

import "math"

func makeGood(s string) string {
	stack := make([]int32, 0)

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		top := stack[len(stack)-1]
		if math.Abs(float64(top-c)) == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	answer := ""
	for _, c := range stack {
		answer += string(c)
	}

	return answer
}
