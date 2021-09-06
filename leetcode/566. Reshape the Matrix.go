package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	resultIterator := 0

	for _, i := range mat {
		for _, j := range i {
			result[resultIterator/c][resultIterator%c] = j
			resultIterator += 1
		}
	}

	return result
}

func main() {
	arr := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
	}

	newArr := matrixReshape(arr, 5, 5)

	for _, i := range newArr {
		for _, j := range i {
			print(j)
			print(" ")
		}

		println()
	}
}
