package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	r, answer := 0, int64(0)
	for l := 0; l < len(nums); l = r {
		for l < len(nums) && (nums[l] < minK || nums[l] > maxK) {
			l++
		}

		if l >= len(nums) {
			return answer
		}

		r = l
		for r < len(nums) && nums[r] >= minK && nums[r] <= maxK {
			r++
		}

		minPos := -1
		maxPos := -1
		for i := l; i < r; i++ {
			if nums[i] == minK {
				minPos = i
			}

			if nums[i] == maxK {
				maxPos = i
			}

			if minPos >= 0 && maxPos >= 0 {
				answer += int64(min(minPos, maxPos) - l + 1)
			}
		}
	}

	return answer
}
