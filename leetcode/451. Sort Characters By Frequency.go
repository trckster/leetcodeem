package main

import (
	"sort"
	"strings"
)

type RuneWithCount struct {
	rune  rune
	count int
}

func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		if cnt, ok := m[r]; ok {
			m[r] = cnt + 1
		} else {
			m[r] = 1
		}
	}

	runes := make([]*RuneWithCount, 0)
	for r, cnt := range m {
		runes = append(runes, &RuneWithCount{r, cnt})
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i].count > runes[j].count
	})

	answer := ""
	for _, r := range runes {
		answer += strings.Repeat(string(r.rune), r.count)
	}
	return answer
}
