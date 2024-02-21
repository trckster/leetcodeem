package main

func getMostSignificantBit(n int) int {
	msb := -1

	for i := 0; i < 64; i++ {
		if 1<<i&n != 0 {
			msb = i
		}
	}

	return msb
}

func rangeBitwiseAnd(left int, right int) int {
	msb := getMostSignificantBit(right)

	answer := 0

	for ; msb >= 0; msb-- {
		lb := 1 << msb & left
		rb := 1 << msb & right

		if lb != rb {
			answer <<= msb + 1
			break
		}

		answer <<= 1
		if lb > 0 {
			answer |= 1
		}
	}

	return answer
}
