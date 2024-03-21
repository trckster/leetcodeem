package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var left *ListNode

	for head.Next != nil {
		right := head.Next
		head.Next = left
		left = head
		head = right
	}

	head.Next = left

	return head
}
