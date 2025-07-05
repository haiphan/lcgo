package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	cur := dummy
	c := 0
	for l1 != nil || l2 != nil || c == 1 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		d := v1 + v2 + c
		c = d / 10
		d = d % 10
		cur.Next = &ListNode{Val: d}
		cur = cur.Next
	}
	return dummy.Next
}
