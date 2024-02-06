package main

import "sort"

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func getSorted(s string) string {
	runes := RuneSlice(s)
	sort.Sort(runes)
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	f := make(map[string][]string)

	for _, s := range strs {
		sorted := getSorted(s)
		if arr, ok := f[sorted]; ok {
			f[sorted] = append(arr, s)
		} else {
			f[sorted] = make([]string, 1)
			f[sorted][0] = s
		}
	}

	answer := make([][]string, 0)
	for _, arr := range f {
		answer = append(answer, arr)
	}
	return answer
}
