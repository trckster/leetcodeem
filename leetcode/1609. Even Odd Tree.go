package main

import aq "github.com/emirpasic/gods/queues/arrayqueue"

func isEvenOddTree(root *TreeNode) bool {
	queue := aq.New()
	queue.Enqueue(root)

	for level := 0; !queue.Empty(); level += 1 {
		item, _ := queue.Dequeue()
		size := queue.Size()

		previousNode := item.(*TreeNode)
		if previousNode.Val%2 == level%2 {
			return false
		}

		if previousNode.Left != nil {
			queue.Enqueue(previousNode.Left)
		}

		if previousNode.Right != nil {
			queue.Enqueue(previousNode.Right)
		}

		for ; size > 0; size-- {
			item, _ = queue.Dequeue()
			node := item.(*TreeNode)

			if level%2 == node.Val%2 {
				return false
			} else if level%2 == 0 && previousNode.Val >= node.Val {
				return false
			} else if level%2 == 1 && previousNode.Val <= node.Val {
				return false
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}

			previousNode = node
		}
	}

	return true
}
