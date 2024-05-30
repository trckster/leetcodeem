package main

func countTriplets(arr []int) int {
	xors := make([][]int, len(arr))
	for i := range xors {
		xors[i] = make([]int, len(arr))
	}

	for i := 0; i < len(arr); i++ {
		xor := 0
		for j := i; j < len(arr); j++ {
			xor ^= arr[j]
			xors[i][j] = xor
		}
	}

	answer := 0
	for i := 0; i < len(arr); i++ {
		for k := i + 1; k < len(arr); k++ {
			for j := i + 1; j <= k; j++ {
				if xors[i][j-1] == xors[j][k] {
					answer++
				}
			}
		}
	}

	return answer
}
