package main

func solveZeroCase(nums []int) int {
	answer := 0

	for i := 0; i < len(nums); i++ {
		cnt := 0
		for ; i < len(nums) && nums[i] == 0; i++ {
			cnt++
		}
		answer += cnt * (cnt + 1) / 2
	}

	return answer
}

func numSubarraysWithSum(nums []int, goal int) int {
	if goal == 0 {
		return solveZeroCase(nums)
	}

	stackedZeros := make([]int, 0)
	for i, n := range nums {
		if n == 1 {
			stackedZeros = append(stackedZeros, -1)
		} else {
			if i > 0 && nums[i-1] == 0 {
				stackedZeros[len(stackedZeros)-1]++
			} else {
				stackedZeros = append(stackedZeros, 1)
			}
		}
	}

	l, r := -1, 0
	for cnt := 0; r < len(stackedZeros); r++ {
		if l == -1 && stackedZeros[r] == -1 {
			l = r
		}

		if stackedZeros[r] == -1 {
			cnt++
		}

		if cnt == goal {
			break
		}
	}

	answer := 0
	for r < len(stackedZeros) {
		println(l, r)
		zerosOnLeft := 1
		if l > 0 && stackedZeros[l-1] != -1 {
			zerosOnLeft += stackedZeros[l-1]
		}
		zerosOnRight := 1
		if r+1 < len(stackedZeros) && stackedZeros[r+1] != -1 {
			zerosOnRight += stackedZeros[r+1]
		}
		answer += zerosOnLeft * zerosOnRight
		r++
		l++
		if r < len(stackedZeros) && stackedZeros[r] != -1 {
			r++
		}
		if l < len(stackedZeros) && stackedZeros[l] != -1 {
			l++
		}
	}
	return answer
}
