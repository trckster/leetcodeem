package main

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var answer *TreeNode
	answerPathDepth := -2
	answerDepth := -2

	var getDepth func(root *TreeNode, currentDepth int) int
	getDepth = func(node *TreeNode, currentDepth int) int {
		if node == nil {
			return -1
		}

		leftDepth := getDepth(node.Left, currentDepth+1)
		rightDepth := getDepth(node.Right, currentDepth+1)
		maxDepth := max(leftDepth, rightDepth)

		thisPathDepth := maxDepth + currentDepth

		if leftDepth == rightDepth && thisPathDepth > answerPathDepth {
			answerPathDepth = thisPathDepth
			answerDepth = currentDepth
			answer = node
		} else if leftDepth == rightDepth && thisPathDepth >= answerPathDepth && currentDepth < answerDepth {
			answerPathDepth = thisPathDepth
			answerDepth = currentDepth
			answer = node
		}

		return maxDepth + 1
	}

	getDepth(root, 0)

	return answer
}
