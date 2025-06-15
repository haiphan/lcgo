package problems

import "lcgo/utils"

type ListNode = utils.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}
