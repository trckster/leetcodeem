package main

func sortColors(nums []int) {
	zeros, ones, twos := 0, 0, 0
	for _, n := range nums {
		if n == 0 {
			zeros++
		} else if n == 1 {
			ones++
		} else {
			twos++
		}
	}

	for i := 0; i < zeros; i++ {
		nums[i] = 0
	}
	for i := zeros; i < zeros+ones; i++ {
		nums[i] = 1
	}
	for i := zeros + ones; i < zeros+ones+twos; i++ {
		nums[i] = 2
	}
}
