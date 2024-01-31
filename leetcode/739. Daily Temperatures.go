package main

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))

	ms := make([][2]int, 0)

	for i, temperature := range temperatures {
		if len(ms) == 0 {
			ms = append(ms, [2]int{i, temperature})
			continue
		}

		if ms[len(ms)-1][1] >= temperature {
			ms = append(ms, [2]int{i, temperature})
			continue
		}

		for len(ms) > 0 && ms[len(ms)-1][1] < temperature {
			answer[ms[len(ms)-1][0]] = i - ms[len(ms)-1][0]
			ms = ms[:len(ms)-1]
		}
		ms = append(ms, [2]int{i, temperature})
	}

	for i := len(ms) - 1; i >= 0; i-- {
		answer[ms[i][0]] = 0
	}

	return answer
}
