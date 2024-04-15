package main

func sumNumbers(root *TreeNode) int {
	answer := 0

	var dfs func(root *TreeNode, beforeSum int)
	dfs = func(root *TreeNode, beforeSum int) {
		updatedSum := beforeSum*10 + root.Val

		if root.Left == nil && root.Right == nil {
			answer += updatedSum
		}

		if root.Left != nil {
			dfs(root.Left, updatedSum)
		}

		if root.Right != nil {
			dfs(root.Right, updatedSum)
		}
	}

	dfs(root, 0)

	return answer
}
