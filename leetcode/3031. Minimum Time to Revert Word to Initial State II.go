package main

func getZ(word string) []int {
	z := make([]int, len(word))
	l, r := 0, 0

	for i := 1; i < len(word); i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}

		for j := i + z[i]; j > r && j < len(word) && word[j] == word[j-i]; j++ {
			z[i] = j - i + 1
		}

		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	return z
}

func minimumTimeToInitialState(word string, k int) int {
	z := getZ(word)

	time := 0
	for i := k; ; i += k {
		time++
		if i >= len(word) || z[i]+i == len(word) {
			break
		}
	}

	return time
}
