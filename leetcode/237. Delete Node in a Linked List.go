package main

func deleteNode(node *ListNode) {
	*node = *node.Next
}
