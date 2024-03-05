package main

func minimumLength(s string) int {
	l, r := 0, len(s)-1

	for l < r {
		leftChar := s[l]
		rightChar := s[r]

		if leftChar != rightChar {
			return r - l + 1
		}

		for ; l < len(s) && s[l] == leftChar; l++ {
		}

		for ; r >= 0 && s[r] == leftChar; r-- {
		}
	}

	return max(r-l+1, 0)
}
