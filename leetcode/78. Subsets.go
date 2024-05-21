package main

func subsets(nums []int) [][]int {
	answer := make([][]int, 1<<len(nums))

	for i := 0; i < 1<<len(nums); i++ {
		answer[i] = make([]int, 0)

		for j, n := range nums {
			if i>>j&1 == 1 {
				answer[i] = append(answer[i], n)
			}
		}
	}

	return answer
}
