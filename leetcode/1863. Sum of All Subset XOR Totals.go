package main

func subsetXORSum(nums []int) int {
	combinations := 1 << len(nums)
	answer := 0

	for i := 1; i < combinations; i++ {
		xor := 0
		for bit, n := range nums {
			if (i>>bit)%2 == 1 {
				xor ^= n
			}
		}
		answer += xor
	}

	return answer
}
