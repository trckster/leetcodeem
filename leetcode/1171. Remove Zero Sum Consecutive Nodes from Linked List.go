package main

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	currentSum := 0

	for i := head; i != nil; i = i.Next {
		currentSum += i.Val
		if currentSum == 0 {
			return removeZeroSumSublists(i.Next)
		}
	}

	head.Next = removeZeroSumSublists(head.Next)
	return head
}
