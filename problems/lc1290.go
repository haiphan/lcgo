package problems

func getDecimalValue(head *ListNode) int {
	num, cur := 0, head
	for cur != nil {
		num = num*2 + cur.Val
		cur = cur.Next
	}
	return num
}
