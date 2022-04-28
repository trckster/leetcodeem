package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	if target <= nums[l] {
		return l
	} else if target == nums[r] {
		return r
	} else if target >= nums[r] {
		return r + 1
	}

	for l+1 != r {
		m := (r-l)/2 + l

		if nums[m] > target {
			r = m
		} else if nums[m] < target {
			l = m
		} else {
			return m
		}
	}

	return r
}

func main() {
	nums1 := []int{1, 3, 5, 6}
	println(searchInsert(nums1, 5))
	println(searchInsert(nums1, 2))
	println(searchInsert(nums1, 7))
	println(searchInsert(nums1, 0))
}
