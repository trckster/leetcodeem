import "math"

func earliestFinishTimeInOrder(firstRideTime []int, firstRideDuration []int, secondRideTime []int, secondRideDuration []int) int {
	bestFinishTime := math.MaxInt32
	for i := range firstRideTime {
		endTime := firstRideTime[i] + firstRideDuration[i]

		if endTime < bestFinishTime {
			bestFinishTime = endTime
		}
	}

	result := math.MaxInt32

	for i := range secondRideTime {
		awaitPenalty := max(secondRideTime[i] - bestFinishTime, 0)
		endTime := bestFinishTime + awaitPenalty + secondRideDuration[i]

		if endTime < result {
			result = endTime
		}
	}

	return result
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	landFirst := earliestFinishTimeInOrder(landStartTime, landDuration, waterStartTime, waterDuration)
	waterFirst := earliestFinishTimeInOrder(waterStartTime, waterDuration, landStartTime, landDuration)

	if landFirst < waterFirst {
		return landFirst
	} else {
		return waterFirst
	}
}
