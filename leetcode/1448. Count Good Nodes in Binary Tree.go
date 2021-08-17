package main

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func countGoodNodes(root *TreeNode, max int) int {
	result := 0

	if root.Val >= max {
		result += 1
	}

	if max < root.Val {
		max = root.Val
	}

	if root.Left != nil {
		result += countGoodNodes(root.Left, max)
	}

	if root.Right != nil {
		result += countGoodNodes(root.Right, max)
	}

	return result
}

func goodNodes(root *TreeNode) int {
	return countGoodNodes(root, -100000)
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	println(goodNodes(root))
}