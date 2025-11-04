package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	dummy := &ListNode{Val: -1}
	cur := dummy
	for _, x := range arr {
		cur.Next = &ListNode{Val: x}
		cur = cur.Next
	}
	return dummy.Next
}

func LinkedListToArray(head *ListNode) []int {
	var arr []int
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur.Val)
	}
	return arr
}
