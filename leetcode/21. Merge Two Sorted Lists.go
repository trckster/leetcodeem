package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(ln *ListNode, val int) *ListNode {
	ln.Next = &ListNode{
		Val:  val,
		Next: nil,
	}

	return ln.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var first int

	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		first = l2.Val
		l2 = l2.Next
	} else if l2 == nil {
		first = l1.Val
		l1 = l1.Next
	} else if l1.Val <= l2.Val {
		first = l1.Val
		l1 = l1.Next
	} else {
		first = l2.Val
		l2 = l2.Next
	}

	head := &ListNode{Val: first}

	current := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current = addNode(current, l1.Val)
			l1 = l1.Next
		} else {
			current = addNode(current, l2.Val)
			l2 = l2.Next
		}
	}

	for l1 != nil {
		current = addNode(current, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		current = addNode(current, l2.Val)
		l2 = l2.Next
	}

	return head
}
