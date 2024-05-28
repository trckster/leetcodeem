package main

func byteDiff(a, b byte) int {
	diff := a - b
	if a < b {
		diff = b - a
	}
	return int(diff)
}

func equalSubstring(s string, t string, maxCost int) int {
	l, r := 0, 0
	answer := 0
	for r < len(s) {
		maxCost -= byteDiff(s[r], t[r])
		r++

		for maxCost < 0 {
			maxCost += byteDiff(s[l], t[l])
			l++
		}

		answer = max(answer, r-l)
	}
	return answer
}
