package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	it := list1
	for i := 0; i < a-1; i++ {
		it = it.Next
	}

	intervalEnd := list1
	for i := 0; i < b; i++ {
		intervalEnd = intervalEnd.Next
	}

	list2end := list2
	for list2end.Next != nil {
		list2end = list2end.Next
	}

	it.Next = list2
	list2end.Next = intervalEnd.Next

	return list1
}
