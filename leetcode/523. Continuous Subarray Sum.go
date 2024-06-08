package main

func checkSubarraySum(nums []int, k int) bool {
	remainders := make(map[int]bool)

	previousSum := 0
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		remainders[previousSum%k] = true
		previousSum = sum
		sum += nums[i]

		if remainders[sum%k] {
			return true
		}
	}

	return false
}
