package main

import "slices"

type Losses map[int]int

func (losses Losses) addWinner(player int) {
	_, ok := losses[player]

	if !ok {
		losses[player] = 0
	}
}

func (losses Losses) addLoser(player int) {
	currentValue, ok := losses[player]

	if !ok {
		losses[player] = 1
	} else {
		losses[player] = currentValue + 1
	}
}

func (losses Losses) getPlayersByLossCount(lossCount int) []int {
	answer := make([]int, 0)

	for player, count := range losses {
		if count == lossCount {
			answer = append(answer, player)
		}
	}

	slices.Sort(answer)

	return answer
}

func findWinners(matches [][]int) [][]int {
	losses := make(Losses)

	for _, match := range matches {
		losses.addWinner(match[0])
		losses.addLoser(match[1])
	}

	answer := make([][]int, 2)
	answer[0] = losses.getPlayersByLossCount(0)
	answer[1] = losses.getPlayersByLossCount(1)

	return answer
}
