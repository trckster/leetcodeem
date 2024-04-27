package main

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	options := make(map[int]int)
	options[0] = 0

	for _, c := range key {
		newOptions := make(map[int]int)

		for k, v := range options {
			i := k
			cnt := 0
			for ring[i] != uint8(c) {
				cnt++
				i--
				if i < 0 {
					i = len(ring) - 1
				}
			}

			if newV, ok := newOptions[i]; ok {
				newOptions[i] = min(newV, v+cnt+1)
			} else {
				newOptions[i] = v + cnt + 1
			}

			i = k
			cnt = 0
			for ; ring[i] != uint8(c); i = (i + 1) % len(ring) {
				cnt++
			}

			if newV, ok := newOptions[i]; ok {
				newOptions[i] = min(newV, v+cnt+1)
			} else {
				newOptions[i] = v + cnt + 1
			}
		}

		options = newOptions
	}

	answer := math.MaxInt

	for _, v := range options {
		if v < answer {
			answer = v
		}
	}

	return answer
}
