package main

func zeroFilledSubarray(nums []int) int64 {
	answer := int64(0)

	sequentialCount := 0
	for _, num := range nums {
		if num == 0 {
			sequentialCount++
			answer += int64(sequentialCount)
		} else {
			sequentialCount = 0
		}
	}

	return answer
}
