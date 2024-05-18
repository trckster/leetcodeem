package main

func countMoves(root *TreeNode) (int, int, int) {
	leftPlus, leftMinus, leftOverall := 0, 0, 0
	if root.Left != nil {
		leftPlus, leftMinus, leftOverall = countMoves(root.Left)
	}

	rightPlus, rightMinus, rightOverall := 0, 0, 0
	if root.Right != nil {
		rightPlus, rightMinus, rightOverall = countMoves(root.Right)
	}

	coinsHere := root.Val + leftPlus + rightPlus - leftMinus - rightMinus

	if coinsHere > 0 {
		return coinsHere - 1, 0, leftOverall + rightOverall + coinsHere - 1
	} else {
		return 0, -coinsHere + 1, leftOverall + rightOverall - coinsHere + 1
	}
}

func distributeCoins(root *TreeNode) int {
	_, _, allMoves := countMoves(root)
	return allMoves
}
