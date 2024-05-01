package main

func reversePrefix(word string, ch byte) string {
	end := 0
	for ; end < len(word) && word[end] != ch; end++ {

	}

	if end == len(word) {
		return word
	}

	answer := ""
	for i := end; i >= 0; i-- {
		answer += string(word[i])
	}
	answer += word[end+1:]

	return answer
}
