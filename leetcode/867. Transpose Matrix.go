package main

func transpose(matrix [][]int) [][]int {
	answer := make([][]int, len(matrix[0]))
	for i, _ := range answer {
		answer[i] = make([]int, len(matrix))
	}

	for i, arr := range matrix {
		for j, value := range arr {
			answer[j][i] = value
		}
	}

	return answer
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	res := transpose(arr)

	for _, arr := range res {
		for _, value := range arr {
			print(value)
			print(" ")
		}

		println()
	}
}
