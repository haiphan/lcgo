package problems

import "slices"

func modifiedList(nums []int, head *ListNode) *ListNode {
	maxv := slices.Max(nums)
	ns := make([]bool, maxv+1)
	for _, x := range nums {
		ns[x] = true
	}
	dummy := &ListNode{Val: -1, Next: head}
	l := dummy
	for r := head; r != nil; r = r.Next {
		if r.Val <= maxv && ns[r.Val] {
			continue
		}
		l = l.Next
		l.Val = r.Val
	}
	l.Next = nil
	return dummy.Next
}
