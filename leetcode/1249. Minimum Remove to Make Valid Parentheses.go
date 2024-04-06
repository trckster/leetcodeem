package main

func minRemoveToMakeValid(s string) string {
	h := 0

	answer := make([]int32, 0)
	for _, c := range s {
		if c == '(' {
			answer = append(answer, c)
			h++
		} else if c == ')' {
			if h > 0 {
				h--
				answer = append(answer, c)
			}
		} else {
			answer = append(answer, c)
		}
	}

	for i := len(answer) - 1; h > 0; i-- {
		if answer[i] == '(' {
			h--
			answer[i] = '-'
		}
	}

	answerAsString := ""
	for _, c := range answer {
		if c != '-' {
			answerAsString += string(c)
		}
	}

	return answerAsString
}
