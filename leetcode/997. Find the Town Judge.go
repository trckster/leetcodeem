package main

type Citizen struct {
	trustedByOthers int
	hasTrusted      bool
}

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}

	trustedCount := make(map[int]Citizen)

	for _, t := range trust {
		trusting, ok1 := trustedCount[t[0]]
		trusted, ok2 := trustedCount[t[1]]

		if ok1 {
			trusting.hasTrusted = true
		} else {
			trusting = Citizen{0, true}
		}

		if ok2 {
			trusted.trustedByOthers += 1
		} else {
			trusted = Citizen{1, false}
		}

		trustedCount[t[0]] = trusting
		trustedCount[t[1]] = trusted
	}

	answer := -1
	for citizenId, citizenData := range trustedCount {
		if !citizenData.hasTrusted && citizenData.trustedByOthers == n-1 {
			if answer != -1 {
				return -1
			}
			answer = citizenId
		}
	}
	return answer
}
