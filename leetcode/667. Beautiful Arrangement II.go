package main

func constructArray(n int, k int) []int {
	result := make([]int, n)

	lower := n - k + 1
	upper := n
	lookingForUpper := true
	for i := 0; i < n; i++ {
		if i + 1 <= n - k {
			result[i] = i + 1
		} else if lookingForUpper {
			result[i] = upper
			upper -= 1
			lookingForUpper = false
		} else {
			result[i] = lower
			lower += 1
			lookingForUpper = true
		}
	}

	return result
}

func main() {
	arr := constructArray(7, 4)

	for _,i := range arr {
		println(i)
	}
}