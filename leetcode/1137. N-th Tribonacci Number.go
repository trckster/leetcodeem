package main

func tribonacci(n int) int {
	if n <= 1 {
		return n
	}

	n1, n2, n3 := 0, 1, 1

	for i := 2; i < n; i++ {
		n1, n2, n3 = n2, n3, n1+n2+n3
	}

	return n3
}
