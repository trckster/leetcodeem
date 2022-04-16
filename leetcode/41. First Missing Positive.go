package main

func firstMissingPositive(nums []int) int {
	for i, _ := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	minimum := 1
	maximum := len(nums)

	for i, _ := range nums {
		number := nums[i]

		for number >= minimum && number <= maximum {
			tmp := nums[number-1]
			nums[number-1] = 0

			number = tmp
		}
	}

	for i, number := range nums {
		if number != 0 {
			return i + 1
		}
	}

	return maximum + 1
}

func main() {
	println(firstMissingPositive([]int{1, 2, 0}))
	println(firstMissingPositive([]int{3, 4, -1, 1}))
	println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	println(firstMissingPositive([]int{7, 8, 9, 0, 2, 3, 1, 11, 12}))
}
