package main

func setNodeOnDepthOnSide(root *TreeNode, val int, depth int, onLeft bool) *TreeNode {
	if depth == 1 {
		newRoot := &TreeNode{Val: val}

		if onLeft {
			newRoot.Left = root
		} else {
			newRoot.Right = root
		}

		return newRoot
	}

	if root == nil {
		return nil
	}

	root.Left = setNodeOnDepthOnSide(root.Left, val, depth-1, true)
	root.Right = setNodeOnDepthOnSide(root.Right, val, depth-1, false)

	return root
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	return setNodeOnDepthOnSide(root, val, depth, true)
}
