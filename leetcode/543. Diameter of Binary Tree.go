package main

func recursiveWayDown(root *TreeNode) (int, int) {
	depth, answer := 1, 0
	leftDepth, leftAnswer := 0, 0
	rightDepth, rightAnswer := 0, 0

	if root.Left != nil {
		leftDepth, leftAnswer = recursiveWayDown(root.Left)
		depth = leftDepth + 1
		answer = max(answer, leftAnswer, depth-1)
	}

	if root.Right != nil {
		rightDepth, rightAnswer = recursiveWayDown(root.Right)
		depth = max(rightDepth+1, depth)
		answer = max(answer, rightAnswer, depth-1)
	}

	answer = max(answer, leftDepth+rightDepth)

	return depth, answer
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, answer := recursiveWayDown(root)
	return answer
}
