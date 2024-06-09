package main

func subarraysDivByK(nums []int, k int) int {
	remainders := make(map[int]int)

	answer := 0
	previousRemainder := 0
	for _, n := range nums {
		remainder := (previousRemainder + n) % k
		if remainder < 0 {
			remainder += k
		} else if remainder == 0 {
			answer++
		}

		answer += remainders[remainder]
		remainders[remainder]++
		previousRemainder = remainder
	}

	return answer
}
