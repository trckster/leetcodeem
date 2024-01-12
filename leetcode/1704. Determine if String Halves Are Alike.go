package main

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func countVowels(s string) int {
	count := 0

	for _, c := range s {
		if isVowel(c) {
			count += 1
		}
	}

	return count
}

func halvesAreAlike(s string) bool {
	return countVowels(s[:len(s)/2]) == countVowels(s[len(s)/2:])
}
