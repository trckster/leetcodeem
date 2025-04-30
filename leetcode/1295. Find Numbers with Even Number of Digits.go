package main

func findNumbers(nums []int) int {
	answer := 0

	for _, num := range nums {
		if 10 <= num && num < 100 || 1000 <= num && num < 10000 || num == 100000 {
			answer++
		}
	}

	return answer
}
