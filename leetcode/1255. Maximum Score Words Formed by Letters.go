package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	wordsAsMap := make([][]int, len(words))
	for i, word := range words {
		wordsAsMap[i] = make([]int, 26)
		for _, c := range word {
			wordsAsMap[i][c-'a']++
		}
	}

	lettersCount := make([]int, 26)
	for _, letter := range letters {
		lettersCount[letter-'a']++
	}

	amountInCurrentWords := make([]int, 26)
	answer := 0
	for i := 1; i < 1<<len(words); i++ {
		for wordIndex := 0; wordIndex < len(wordsAsMap); wordIndex++ {
			if i&(1<<wordIndex) > 0 {
				for j, c := range wordsAsMap[wordIndex] {
					amountInCurrentWords[j] += c
				}
			}
		}

		newAnswer := 0
		for j, charsCount := range amountInCurrentWords {
			if lettersCount[j] < charsCount {
				newAnswer = 0
				break
			}

			newAnswer += score[j] * charsCount
		}

		answer = max(answer, newAnswer)

		for j := range amountInCurrentWords {
			amountInCurrentWords[j] = 0
		}
	}

	return answer
}
