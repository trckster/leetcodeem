package main

func minOperations(nums []int, k int) int {
	for _, n := range nums {
		k ^= n
	}
	answer := 0
	for k > 0 {
		answer += k % 2
		k >>= 1
	}
	return answer
}
