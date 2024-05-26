package main

const mod = 1000000007

func checkRecord(n int) int {
	lastRecord := make([]int, 6)
	lastRecord[0] = 3 // Answer
	lastRecord[1] = 1 // +A
	lastRecord[2] = 1 // +A -L
	lastRecord[3] = 0 // +A +L
	lastRecord[4] = 1 // -A -L
	lastRecord[5] = 1 // -A +L

	currentRecord := make([]int, 6)
	for i := 1; i < n; i++ {
		currentRecord[0] = lastRecord[0]*2 - lastRecord[1] + lastRecord[2] + lastRecord[3] + lastRecord[4] + lastRecord[5]
		currentRecord[1] = lastRecord[2] + lastRecord[3] + lastRecord[0]
		currentRecord[2] = lastRecord[0]
		currentRecord[3] = lastRecord[2]
		currentRecord[4] = lastRecord[0] - lastRecord[1]
		currentRecord[5] = lastRecord[4]

		for j, v := range currentRecord {
			lastRecord[j] = v % mod
			if lastRecord[j] < 0 {
				lastRecord[j] = mod + lastRecord[j]
			}
		}
	}

	return lastRecord[0]
}
