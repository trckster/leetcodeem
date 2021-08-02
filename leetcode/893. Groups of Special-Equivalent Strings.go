package main

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	arr := strings.Split(s, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}

func smartSort(s string) string {
	even := ""
	odd := ""

	for i, r := range s {
		c := string(r)

		if i%2 == 0 {
			even += c
		} else {
			odd += c
		}
	}

	even = sortString(even)
	odd = sortString(odd)

	return even + odd
}

func numSpecialEquivGroups(words []string) int {
	set := make(map[string]bool)

	for i, word := range words {
		words[i] = smartSort(word)
		set[words[i]] = true
	}

	return len(set)
}

func main() {
	test := []string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}
	answer := numSpecialEquivGroups(test)
	println(answer)
}
