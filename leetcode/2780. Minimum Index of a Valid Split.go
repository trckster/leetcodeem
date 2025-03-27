package main

func minimumIndex(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	dominant := nums[0]
	dominantCount := 1
	for num, count := range counts {
		if count > dominantCount {
			dominantCount = count
			dominant = num
		}
	}

	left := 0
	right := dominantCount
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == dominant {
			left++
			right--
		}

		if (i+1)/2 < left && (len(nums)-i-1)/2 < right {
			return i
		}
	}

	return -1
}
