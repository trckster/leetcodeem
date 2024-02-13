package main

func firstPalindrome(words []string) string {
	for _, s := range words {
		palindrome := true

		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-i-1] {
				palindrome = false
			}
		}

		if palindrome {
			return s
		}
	}
	return ""
}
