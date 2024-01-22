package main

func findErrorNums(nums []int) []int {
	track := make([]bool, len(nums))
	answer := make([]int, 2)

	for _, number := range nums {
		if track[number-1] {
			answer[0] = number
		}

		track[number-1] = true
	}

	for i := 0; i < len(nums); i++ {
		if !track[i] {
			answer[1] = i + 1
		}
	}

	return answer
}
