package main

func smallestFromLeaf(root *TreeNode) string {
	answer := ""

	var dfs func(root *TreeNode, suffix string)
	dfs = func(root *TreeNode, suffix string) {
		s := string(rune(97+root.Val)) + suffix

		if root.Left == nil && root.Right == nil {
			if answer == "" || answer > s {
				answer = s
			}
		}

		if root.Left != nil {
			dfs(root.Left, s)
		}

		if root.Right != nil {
			dfs(root.Right, s)
		}
	}
	dfs(root, "")

	return answer
}
