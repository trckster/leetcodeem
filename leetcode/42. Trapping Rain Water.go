package main

func trap(height []int) int {
	stack := make([][]int, 0)

	answer := 0
	for i, h := range height {
		if len(stack) == 0 {
			stack = append(stack, []int{h, i})
			continue
		}

		currentWaterHeight := stack[len(stack)-1][1]
		for len(stack) > 0 && stack[len(stack)-1][0] < h {
			leftPoint := stack[len(stack)-1]

			trapHeight := min(leftPoint[0], h)
			additionalWaterHeight := trapHeight - currentWaterHeight
			trapWidth := i - leftPoint[1] - 1

			answer += trapWidth * additionalWaterHeight

			currentWaterHeight = trapHeight

			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			leftPoint := stack[len(stack)-1]

			trapHeight := min(leftPoint[0], h)
			additionalWaterHeight := trapHeight - currentWaterHeight
			trapWidth := i - leftPoint[1] - 1

			answer += trapWidth * additionalWaterHeight

			currentWaterHeight = trapHeight
		}

		stack = append(stack, []int{h, i})
	}

	return answer
}
