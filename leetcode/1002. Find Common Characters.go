package main

func commonChars(words []string) []string {
	wm := make([]map[rune]int, len(words))
	for i, word := range words {
		wm[i] = make(map[rune]int)
		for _, ch := range word {
			wm[i][ch]++
		}
	}

	result := make([]string, 0)
	for k := range wm[0] {
		letterCount := wm[0][k]

		for i := 1; i < len(words); i++ {
			letterCount = min(letterCount, wm[i][k])
		}

		for i := 0; i < letterCount; i++ {
			result = append(result, string(k))
		}
	}

	return result
}
