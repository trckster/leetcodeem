package main

func evaluateTree(root *TreeNode) bool {
	if root.Val == 1 {
		return true
	}

	if root.Val == 0 {
		return false
	}

	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}

	return evaluateTree(root.Left) && evaluateTree(root.Right)
}
