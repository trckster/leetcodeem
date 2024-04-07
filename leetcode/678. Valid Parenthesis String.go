package main

func checkValidString(s string) bool {
	l, r := 0, 0

	for _, c := range s {
		if c == '(' {
			l++
			r++
		} else if c == ')' {
			l = max(l-1, 0)
			r--
		} else {
			r++
			l = max(l-1, 0)
		}

		if r < 0 {
			return false
		}
	}

	return l == 0
}
