package main

import "strings"

func customSortString(order string, s string) string {
	orderMap := make(map[rune]bool)
	for _, c := range order {
		orderMap[c] = true
	}

	sCount := make(map[rune]int)
	remainder := ""
	for _, c := range s {
		if orderMap[c] {
			sCount[c]++
		} else {
			remainder += string(c)
		}
	}

	answer := ""
	for _, c := range order {
		answer += strings.Repeat(string(c), sCount[c])
	}

	return answer + remainder
}
