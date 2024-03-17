package main

func insert(intervals [][]int, newInterval []int) [][]int {
	answer := make([][]int, 0)

	added := false
	start := -1
	end := -1

	for _, interval := range intervals {
		if added || interval[1] < newInterval[0] {
			answer = append(answer, interval)
			continue
		}

		if newInterval[1] < interval[0] {
			if start == -1 {
				start = newInterval[0]
			}
			if end == -1 {
				end = newInterval[1]
			}
			answer = append(answer, []int{start, end})
			added = true
			answer = append(answer, interval)
			continue
		}

		if newInterval[0] <= interval[1] && start == -1 {
			start = min(interval[0], newInterval[0])
		}

		end = max(end, interval[1], newInterval[1])
	}

	if !added {
		if start == -1 {
			start = newInterval[0]
		}
		if end == -1 {
			end = newInterval[1]
		}

		answer = append(answer, []int{start, end})
	}

	return answer
}
