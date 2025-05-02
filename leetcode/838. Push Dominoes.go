package main

import aq "github.com/emirpasic/gods/queues/arrayqueue"

type Fall struct {
	i         int
	direction rune
}

func pushDominoes(dominoes string) string {
	state := []rune(dominoes)
	falls := aq.New()

	for i, d := range dominoes {
		if d == '.' {
			continue
		}

		if d == 'R' && i+1 < len(dominoes) && dominoes[i+1] == '.' {
			falls.Enqueue(Fall{i, d})
		}

		if d == 'L' && i-1 >= 0 && dominoes[i-1] == '.' {
			falls.Enqueue(Fall{i, d})
		}
	}

	for !falls.Empty() {
		item, _ := falls.Dequeue()
		move := item.(Fall)

		if move.direction == 'R' {
			if move.i+1 >= len(state) {
				continue
			}

			if state[move.i+1] != '.' {
				continue
			}

			if move.i+2 == len(state) {
				state[move.i+1] = 'R'
				continue
			}

			if state[move.i+2] == 'L' {
				state[move.i+1] = 'x'
				continue
			}

			state[move.i+1] = 'R'
			falls.Enqueue(Fall{move.i + 1, 'R'})
		}

		if move.direction == 'L' {
			if move.i-1 < 0 {
				continue
			}

			if state[move.i-1] != '.' {
				continue
			}

			state[move.i-1] = 'L'
			falls.Enqueue(Fall{move.i - 1, 'L'})
		}
	}

	for i, s := range state {
		if s == 'x' {
			state[i] = '.'
		}
	}

	return string(state)
}
