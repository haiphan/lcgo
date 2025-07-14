package problems

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		node := &ListNode{Val: gcd(cur.Val, next.Val), Next: next}
		cur.Next = node
		cur = next
	}
	return head
}
