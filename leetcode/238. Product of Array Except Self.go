package main

func productExceptSelf(nums []int) []int {
	zerosCount := 0
	mult := 1
	for _, n := range nums {
		if n == 0 {
			zerosCount++
		} else {
			mult *= n
		}
	}

	if zerosCount > 1 {
		for i := range nums {
			nums[i] = 0
		}
		return nums
	} else if zerosCount == 1 {
		for i, n := range nums {
			nums[i] = 0
			if n == 0 {
				nums[i] = mult
			}
		}
		return nums
	}

	for i, n := range nums {
		nums[i] = mult / n
	}
	return nums
}
