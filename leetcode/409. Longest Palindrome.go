package main

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	answer := 0
	hasOdd := false
	for _, v := range m {
		if v%2 == 0 {
			answer += v
		} else {
			answer += v - 1
			hasOdd = true
		}
	}

	if hasOdd {
		answer++
	}

	return answer
}
