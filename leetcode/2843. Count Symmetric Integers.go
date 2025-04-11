package main

func isSymmetric(n int) bool {
	if n < 10 || n >= 100 && n < 1000 {
		return false
	}

	if n < 100 {
		return n%10 == n/10
	}

	left := n/1000 + (n/100)%10
	right := (n%100)/10 + n%10

	return left == right
}

func countSymmetricIntegers(low int, high int) int {
	answer := 0

	for i := low; i <= high; i++ {
		if isSymmetric(i) {
			answer++
		}
	}

	return answer
}
