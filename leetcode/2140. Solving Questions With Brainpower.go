package main

func mostPoints(questions [][]int) int64 {
	points := make([]int64, len(questions)+1)

	for i := 0; i < len(questions); i++ {
		if i > 0 && points[i-1] > points[i] {
			points[i] = points[i-1]
		}

		newPoints := points[i] + int64(questions[i][0])

		nextAvailableAt := i + questions[i][1] + 1
		if nextAvailableAt > len(questions) {
			nextAvailableAt = len(questions)
		}

		points[nextAvailableAt] = max(points[nextAvailableAt], newPoints)
	}

	return points[len(questions)]
}
