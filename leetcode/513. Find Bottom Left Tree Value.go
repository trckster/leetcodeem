package main

import "errors"

func getDepth(root *TreeNode) int {
	depth := 1

	if root.Left != nil {
		depth = max(depth, getDepth(root.Left)+1)
	}

	if root.Right != nil {
		depth = max(depth, getDepth(root.Right)+1)
	}

	return depth
}

func getLeftmostElementOnLevel(root *TreeNode, currentDepth int, depth int) (int, error) {
	currentDepth += 1

	if depth == currentDepth {
		return root.Val, nil
	}

	if root.Left != nil {
		if result, err := getLeftmostElementOnLevel(root.Left, currentDepth, depth); err == nil {
			return result, nil
		}
	}

	if root.Right != nil {
		if result, err := getLeftmostElementOnLevel(root.Right, currentDepth, depth); err == nil {
			return result, nil
		}
	}

	return 0, errors.New("no deep elements on this branch")
}

func findBottomLeftValue(root *TreeNode) int {
	depth := getDepth(root)

	answer, _ := getLeftmostElementOnLevel(root, 0, depth)

	return answer
}
