package main

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func replaceNonCoprimes(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = nums[0]
	ai := 0

	for _, n := range nums[1:] {
		d := gcd(n, answer[ai])

		if d <= 1 {
			ai += 1
			answer[ai] = n

			continue
		}

		answer[ai] = answer[ai] / d * n
		for ai > 0 {
			d = gcd(answer[ai], answer[ai-1])

			if d <= 1 {
				break
			}

			answer[ai-1] = answer[ai-1] / d * answer[ai]
			ai--
		}
	}

	return answer[:ai+1]
}
