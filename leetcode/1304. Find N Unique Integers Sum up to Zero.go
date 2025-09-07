package main

func sumZero(n int) []int {
	res := make([]int, n)

	for i := 0; i < n/2; i++ {
		res[i*2+1] = -(n / 2) + i
	}
	for i := 0; i < n/2; i++ {
		res[i*2] = n/2 - i
	}
	if n%2 != 0 {
		res[n-1] = 0
	}

	return res
}
