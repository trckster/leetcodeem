package main

import (
	"strconv"
	"strings"
)

func asIntArray(v string) []int {
	s := strings.Split(v, ".")
	numbers := make([]int, len(s))

	for i, n := range s {
		numbers[i], _ = strconv.Atoi(n)
	}

	return numbers
}

func compareVersion(version1 string, version2 string) int {
	a := asIntArray(version1)
	b := asIntArray(version2)

	for i := 0; i < len(a) || i < len(b); i++ {
		left := 0
		if i < len(a) {
			left = a[i]
		}

		right := 0
		if i < len(b) {
			right = b[i]
		}

		if left < right {
			return -1
		} else if right < left {
			return 1
		}
	}

	return 0
}
