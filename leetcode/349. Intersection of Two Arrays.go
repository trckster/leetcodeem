package main

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for _, n := range nums1 {
		m1[n] = true
	}
	for _, n := range nums2 {
		m2[n] = true
	}

	answer := make([]int, 0)
	for n1 := range m1 {
		if m2[n1] {
			answer = append(answer, n1)
		}
	}

	return answer
}
