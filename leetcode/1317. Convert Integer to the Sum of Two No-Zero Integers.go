package main

func hasZero(n int) bool {
	for {
		if n == 0 {
			return false
		}

		if n%10 == 0 {
			return true
		}

		n /= 10
	}
}

func getNoZeroIntegers(n int) []int {
	a := n / 2
	b := n/2 + n%2

	for {
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}

		a--
		b++
	}
}
