package problems

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := (len(nums) - 1) >> 1
	return &TreeNode{Val: nums[m], Left: sortedArrayToBST(nums[:m]), Right: sortedArrayToBST(nums[m+1:])}
}
