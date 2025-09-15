package main

func canBeTypedWords(text string, brokenLetters string) int {
	m := make(map[byte]bool)

	for _, c := range brokenLetters {
		m[byte(c)] = true
	}

	result := 0

	for i := 0; i < len(text); i++ {
		brokenWord := false
		for i < len(text) {
			if text[i] == ' ' {
				break
			}

			if m[text[i]] {
				brokenWord = true
			}

			i++
		}

		if !brokenWord {
			result++
		}
	}

	return result
}
