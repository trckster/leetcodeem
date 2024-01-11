package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func walk(root *TreeNode) (int, int, int) {
	answer := -1
	minVal := root.Val
	maxVal := root.Val

	if root.Left != nil {
		leftAnswer, leftMin, leftMax := walk(root.Left)

		if leftMin < minVal {
			minVal = leftMin
		}

		if leftMax > maxVal {
			maxVal = leftMax
		}

		if leftAnswer > answer {
			answer = leftAnswer
		}
	}

	if root.Right != nil {
		rightAnswer, rightMin, rightMax := walk(root.Right)

		if rightMin < minVal {
			minVal = rightMin
		}

		if rightMax > maxVal {
			maxVal = rightMax
		}

		if rightAnswer > answer {
			answer = rightAnswer
		}
	}

	newDiff := max(abs(root.Val-minVal), abs(root.Val-maxVal))

	return max(answer, newDiff), minVal, maxVal
}

func maxAncestorDiff(root *TreeNode) int {
	answer, _, _ := walk(root)
	return answer
}
