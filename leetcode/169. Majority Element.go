package main

func majorityElement(nums []int) int {
	answer := 0
	answerCount := 0

	for _, n := range nums {
		if answerCount == 0 {
			answer = n
			answerCount += 1
		} else if answer == n {
			answerCount += 1
		} else {
			answerCount -= 1
		}
	}

	return answer
}
