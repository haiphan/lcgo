package problems

type NumArray struct {
	prefix []int
}

func RangeSumConstructor(nums []int) NumArray {
	N := len(nums)
	prefix := make([]int, N+1, N+1)
	for i := range N {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
