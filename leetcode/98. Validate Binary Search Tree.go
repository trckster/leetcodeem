package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode, min int, max int) bool {
	if root.Val >= max || root.Val <= min {
		return false
	}

	if root.Left != nil {
		if !solve(root.Left, min, root.Val) {
			return false
		}
	}

	if root.Right != nil {
		if !solve(root.Right, root.Val, max) {
			return false
		}
	}

	return true
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if !solve(root.Left, math.MinInt64, root.Val) {
			return false
		}
	}

	if root.Right != nil {
		if !solve(root.Right, root.Val, math.MaxInt64) {
			return false
		}
	}

	return true
}

func main() {
	root := &TreeNode{
		Val: -2147483648,
		Right: &TreeNode{
			Val: 2147483647,
		},
	}

	if isValidBST(root) {
		println("Valid")
	} else {
		println("Invalid")
	}
}
