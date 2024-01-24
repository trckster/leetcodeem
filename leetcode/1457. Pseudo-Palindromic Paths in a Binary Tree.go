package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type DigitsCount struct {
	data map[int]int
	odds int
	size int
}

var digitsCount DigitsCount

func (cnt *DigitsCount) bump(digit int) {
	cnt.size += 1
	if _, ok := cnt.data[digit]; !ok {
		cnt.data[digit] = 1
		cnt.odds += 1
	} else {
		cnt.data[digit] += 1
		if cnt.data[digit]%2 == 1 {
			cnt.odds += 1
		} else {
			cnt.odds -= 1
		}
	}
}

func (cnt *DigitsCount) dump(digit int) {
	cnt.size -= 1
	cnt.data[digit] -= 1
	if cnt.data[digit]%2 == 1 {
		cnt.odds += 1
	} else {
		cnt.odds -= 1
	}
}

func walk(root *TreeNode) int {
	answer := 0
	digitsCount.bump(root.Val)

	if root.Left == nil && root.Right == nil {
		if digitsCount.odds == 0 || digitsCount.odds == 1 && digitsCount.size%2 == 1 {
			answer += 1
		}
	}

	if root.Left != nil {
		answer += walk(root.Left)
	}

	if root.Right != nil {
		answer += walk(root.Right)
	}

	digitsCount.dump(root.Val)
	return answer
}

func pseudoPalindromicPaths(root *TreeNode) int {
	digitsCount = DigitsCount{
		make(map[int]int), 0, 0,
	}

	return walk(root)
}
