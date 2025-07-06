package problems

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	nc2   map[int]int
}

func FindSumPairsConstructor(nums1 []int, nums2 []int) FindSumPairs {
	nc2 := make(map[int]int)
	for _, x := range nums2 {
		nc2[x]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, nc2: nc2}
}

func (this *FindSumPairs) Add(index int, val int) {
	ov := this.nums2[index]
	this.nums2[index] += val
	this.nc2[ov]--
	this.nc2[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	cnt := 0
	for _, x := range this.nums1 {
		if v, has := this.nc2[tot-x]; has {
			cnt += v
		}
	}
	return cnt
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
