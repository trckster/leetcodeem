package main

func maximumTripletValue(nums []int) int64 {
	maxValue := max(nums[0], nums[1])
	maxDiff := nums[0] - nums[1]
	answer := int64(0)

	for k := 2; k < len(nums); k++ {
		answer = max(answer, int64(maxDiff)*int64(nums[k]))

		if nums[k] > maxValue {
			maxValue = nums[k]
		} else if maxValue-nums[k] > maxDiff {
			maxDiff = maxValue - nums[k]
		}
	}

	return answer
}
