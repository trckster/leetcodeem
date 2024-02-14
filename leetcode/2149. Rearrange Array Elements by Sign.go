package main

func rearrangeArray(nums []int) []int {
	posI := 0
	negI := 0
	answer := make([]int, len(nums))

	for i, _ := range nums {
		needPos := i == 0 || answer[i-1] < 0

		if needPos {
			for nums[posI] < 0 {
				posI++
			}
			answer[i] = nums[posI]
			posI++
		} else {
			for nums[negI] > 0 {
				negI++
			}
			answer[i] = nums[negI]
			negI++
		}
	}

	return answer
}
