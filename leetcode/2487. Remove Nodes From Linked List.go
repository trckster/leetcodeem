package main

func removeNodesAndReturnMax(head *ListNode) (*ListNode, int) {
	if head == nil {
		return nil, 0
	}

	next, maxValue := removeNodesAndReturnMax(head.Next)

	if head.Val < maxValue {
		return next, maxValue
	}

	head.Next = next

	return head, head.Val
}

func removeNodes(head *ListNode) *ListNode {
	result, _ := removeNodesAndReturnMax(head)

	return result
}
