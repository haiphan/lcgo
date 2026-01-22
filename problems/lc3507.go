package problems

func minimumPairRemoval(nums []int) int {
	BIG := 60000
	dummy := &ListNode{Val: -1}
	ans := 0
	cur := dummy
	for _, x := range nums {
		node := &ListNode{Val: x}
		cur.Next = node
		cur = node
	}
	head := dummy.Next
	cur = head
	for cur.Next != nil {
		if cur.Next.Val >= cur.Val {
			cur = cur.Next
		} else {
			cur = head
			left := cur
			minSum := BIG
			for cur.Next != nil {
				sum := cur.Val + cur.Next.Val
				if sum < minSum {
					minSum = sum
					left = cur
				}
				cur = cur.Next
			}
			right := left.Next
			left.Val += right.Val
			left.Next = right.Next
			right.Next = nil
			ans++
			cur = head
		}
	}

	return ans
}
