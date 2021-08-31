package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var newHead, runner *ListNode

	for ; head != nil; head = head.Next {
		if head.Val != val {
			if newHead == nil {
				newHead = &ListNode{
					Val: head.Val,
				}
				runner = newHead
			} else {
				runner.Next = &ListNode{
					Val: head.Val,
				}
				runner = runner.Next
			}
		}
	}

	return newHead
}

func outputList(head *ListNode) {
	println("-------")

	if head == nil {
		println("Empty list")
	}

	for ; head != nil; head = head.Next {
		println("List element:", head.Val)
	}
}

func main() {
	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  12,
			Next: &ListNode{Val: 5},
		},
	}

	outputList(head)

	head = removeElements(head, 5)

	outputList(head)

	head = removeElements(head, 12)

	outputList(head)
}
