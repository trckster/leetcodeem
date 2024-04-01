package main

import "strings"

func lengthOfLastWord(s string) int {
	parts := strings.Split(s, " ")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] != "" {
			return len(parts[i])
		}
	}
	return 0
}
