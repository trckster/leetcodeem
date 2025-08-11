package main

func modp(n int, p int, mod int) int {
	answer := int64(1)
	for ; p > 0; p-- {
		answer *= int64(n)
		answer %= int64(mod)
	}
	return int(answer)
}

func buildPowers(n int) []int {
	powers := make([]int, 0)
	p := 0

	for n > 0 {
		if n&1 == 1 {
			powers = append(powers, p)
		}

		n >>= 1
		p++
	}

	return powers
}

func productQueries(n int, queries [][]int) []int {
	powers := buildPowers(n)
	sums := make([]int, len(powers))

	sums[0] = powers[0]
	for i, p := range powers[1:] {
		sums[i+1] = sums[i] + p
	}

	answer := make([]int, len(queries))
	for i, q := range queries {
		minus := 0
		if q[0] > 0 {
			minus = sums[q[0]-1]
		}
		power := sums[q[1]] - minus
		answer[i] = modp(2, power, 1000000007)
	}

	return answer
}
