package main

func sumOfLeftLeaves(root *TreeNode) int {
	answer := 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			answer += root.Left.Val
		} else {
			answer += sumOfLeftLeaves(root.Left)
		}
	}

	if root.Right != nil {
		answer += sumOfLeftLeaves(root.Right)
	}

	return answer
}
