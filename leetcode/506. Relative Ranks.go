package main

import (
	"sort"
	"strconv"
)

type Athlete struct {
	id       int
	score    int
	position string
}

func findRelativeRanks(score []int) []string {
	athletes := make([]Athlete, len(score))

	for i, s := range score {
		athletes[i] = Athlete{i, s, ""}
	}

	sort.Slice(athletes, func(i, j int) bool {
		return athletes[i].score > athletes[j].score
	})

	for i := range athletes {
		switch i {
		case 0:
			athletes[i].position = "Gold Medal"
		case 1:
			athletes[i].position = "Silver Medal"
		case 2:
			athletes[i].position = "Bronze Medal"
		default:
			athletes[i].position = strconv.Itoa(i + 1)
		}
	}

	answer := make([]string, len(score))

	sort.Slice(athletes, func(i, j int) bool {
		return athletes[i].id < athletes[j].id
	})

	for i, a := range athletes {
		answer[i] = a.position
	}

	return answer
}
