package main

func numSteps(s string) int {
	steps := 0

	bit := len(s) - 1
	for ; s[bit] == '0'; bit-- {
		steps++
	}

	if bit == 0 {
		return steps
	}

	steps++
	for ; bit >= 0; bit-- {
		if s[bit] == '0' {
			steps += 2
		} else {
			steps += 1
		}
	}

	return steps
}
