package main

func reverse(head *ListNode) *ListNode {
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

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	length := 0
	for it := head; it != nil; it = it.Next {
		length++
	}

	second := head
	for i := 0; i < length/2; i++ {
		second = second.Next
	}

	second = reverse(second)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for head != nil {
			ch1 <- head.Val
			head = head.Next
		}
		close(ch1)
	}()
	go func() {
		for second != nil {
			ch2 <- second.Val
			second = second.Next
		}
		close(ch2)
	}()

	for {
		a := <-ch1
		b, ok2 := <-ch2

		if !ok2 {
			return true
		}

		if a != b {
			return false
		}
	}
}
