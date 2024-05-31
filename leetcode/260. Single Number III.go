package main

func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	saveXor := xor

	firstDifferingBit := 0
	for xor != 1 && xor != -1 {
		xor >>= 1
		firstDifferingBit++
	}

	group1 := 1 << firstDifferingBit
	group1xor := 0
	for _, n := range nums {
		if n&group1 != 0 {
			group1xor ^= n
		}
	}

	return []int{group1xor, group1xor ^ saveXor}
}
