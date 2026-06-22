func maxNumberOfBalloons(text string) int {
	m := map[rune]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}

	for _, c := range text {
		if v, ok := m[c]; ok {
			m[c] = v + 1
		}
	}

	m['l'] /= 2
	m['o'] /= 2

	min := len(text)
	for _, v := range m {
		if v < min {
			min = v
		}
	}

	return min
}
