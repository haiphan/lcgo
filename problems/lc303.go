package problems

type NumArray struct {
	prefix []int
}

func RangeSumConstructor(nums []int) NumArray {
	N := len(nums)
	prefix := make([]int, N+1)
	for i := range N {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{prefix: prefix}
}

func (sr *NumArray) SumRange(left int, right int) int {
	return sr.prefix[right+1] - sr.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
