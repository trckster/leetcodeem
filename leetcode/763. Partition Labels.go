package main

func partitionLabels(s string) []int {
	chars := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}

	charsCount := 0
	charsToTheLeft := make(map[byte]int)

	answer := make([]int, 0)
	currentLength := 0
	for i := 0; i < len(s); i++ {
		currentLength++

		if charsToTheLeft[s[i]] == 0 {
			charsToTheLeft[s[i]] = chars[s[i]]
			charsCount++
		}

		charsToTheLeft[s[i]]--
		if charsToTheLeft[s[i]] == 0 {
			charsCount--
		}

		if charsCount == 0 {
			answer = append(answer, currentLength)
			currentLength = 0
		}
	}

	return answer
}
