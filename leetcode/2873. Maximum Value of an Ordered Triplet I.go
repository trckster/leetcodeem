package main

func maximumTripletValue(nums []int) int64 {
	result := int64(0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				tripletValue := (nums[i] - nums[j]) * nums[k]
				result = max(result, int64(tripletValue))
			}
		}
	}

	return result
}
