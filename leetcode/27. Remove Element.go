package main

func removeElement(nums []int, val int) int {
	realI := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}

		nums[realI] = nums[i]
		realI++
	}

	return realI
}
