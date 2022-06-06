package main

func runRecursion(node *ListNode, c chan *ListNode) {
	if node.Next != nil {
		runRecursion(node.Next, c)
	}

	c <- node
}

func collectNodes(node *ListNode, c chan *ListNode) {
	runRecursion(node, c)
	close(c)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodesA := make(chan *ListNode)
	nodesB := make(chan *ListNode)

	go collectNodes(headA, nodesA)
	go collectNodes(headB, nodesB)

	var answer *ListNode

	for {
		valA, okA := <-nodesA
		valB, okB := <-nodesB

		if !okA || !okB {
			return answer
		}

		if valA == valB {
			answer = valA
		} else {
			return answer
		}
	}
}
