package main

func double(head *ListNode) int {
	if head == nil {
		return 0
	}

	addition := double(head.Next)

	head.Val = head.Val*2 + addition

	if head.Val >= 10 {
		head.Val -= 10
		return 1
	}

	return 0
}

func doubleIt(head *ListNode) *ListNode {
	addition := double(head)

	if addition == 0 {
		return head
	}

	return &ListNode{addition, head}
}
