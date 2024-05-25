package main

import "strings"

func wordBreak(s string, wordDict []string) []string {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}

	answer := make([]string, 0)
	for i := 0; i < 1<<(len(s)-1); i++ {
		newLine := make([]string, 0)

		from := 0
		for j := 0; j < len(s); j++ {
			if i&(1<<j) > 0 {
				anotherWord := s[from : j+1]
				from = j + 1

				newLine = append(newLine, anotherWord)
			}
		}
		newLine = append(newLine, s[from:])

		hasOnlyValidWords := true
		for _, word := range newLine {
			if !set[word] {
				hasOnlyValidWords = false
				break
			}
		}

		if hasOnlyValidWords {
			answer = append(answer, strings.Join(newLine, " "))
		}
	}

	return answer
}
