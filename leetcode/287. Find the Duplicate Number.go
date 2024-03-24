package main

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	fast = nums[0]

	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}
