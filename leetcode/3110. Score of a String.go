package main

func scoreOfString(s string) int {
	answer := 0

	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff >= 0 {
			answer += diff
		} else {
			answer -= diff
		}
	}

	return answer
}
