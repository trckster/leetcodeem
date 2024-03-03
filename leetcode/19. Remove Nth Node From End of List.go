package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	for iterator := head; iterator != nil; iterator = iterator.Next {
		size++
	}

	fromStart := size - n - 1
	if fromStart < 0 {
		return head.Next
	}

	iterator := head
	for ; fromStart > 0; fromStart-- {
		iterator = iterator.Next
	}

	iterator.Next = iterator.Next.Next

	return head
}
