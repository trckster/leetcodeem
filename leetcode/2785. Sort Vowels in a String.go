package main

import (
	"sort"
	"strings"
)

func sortVowels(s string) string {
	vowels := make(map[rune]bool)
	vowels['a'] = true
	vowels['e'] = true
	vowels['i'] = true
	vowels['o'] = true
	vowels['u'] = true

	result := make([]rune, len(s))

	v := make([]rune, 0)

	for i, ch := range s {
		chLower := rune(strings.ToLower(string(ch))[0])

		if vowels[chLower] {
			v = append(v, ch)
		} else {
			result[i] = ch
		}
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})

	vi := 0
	for i, ch := range result {
		if ch == 0 {
			result[i] = v[vi]
			vi++
		}
	}

	return string(result)
}
