package main

func appendCharacters(s string, t string) int {
	ti := 0
	matched := 0
	for i := range s {
		if s[i] == t[ti] {
			matched++
			ti++
		}

		if ti >= len(t) {
			break
		}
	}

	return len(t) - matched
}
