package main

func findDuplicates(nums []int) []int {
	var updateCell func(n int)
	updateCell = func(n int) {
		occupiedBy := nums[n-1]

		if occupiedBy > 0 {
			nums[n-1] = -1
			updateCell(occupiedBy)
		} else if occupiedBy == 0 {
			nums[n-1] = -1
		} else {
			nums[n-1] = -2
		}
	}

	for i, n := range nums {
		if n < 0 {
			continue
		}

		nums[i] = 0
		updateCell(n)
	}

	answer := make([]int, 0)
	for i, n := range nums {
		if n == -2 {
			answer = append(answer, i+1)
		}
	}

	return answer
}
