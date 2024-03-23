package main

func reorderList(head *ListNode) {
	arr := make([]*ListNode, 0)

	for it := head; it != nil; it = it.Next {
		arr = append(arr, it)
	}

	reorderedArr := make([]*ListNode, len(arr))
	for i, j := 0, 0; i < len(arr); i, j = i+2, j+1 {
		reorderedArr[i] = arr[j]

		if j >= len(arr)-j-1 {
			break
		}

		reorderedArr[i+1] = arr[len(arr)-j-1]
	}

	it := head
	for i := 1; i < len(reorderedArr); i++ {
		it.Next = reorderedArr[i]
		it = it.Next
	}

	it.Next = nil
}
