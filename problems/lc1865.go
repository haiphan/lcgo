package problems

type FindSumPairs struct {
	nums1    []int
	nums2    []int
	countMap map[int]int
}

func FindSumPairsConstructor(nums1 []int, nums2 []int) FindSumPairs {
	nc2 := make(map[int]int)
	for _, x := range nums2 {
		nc2[x]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, countMap: nc2}
}

func (findNP *FindSumPairs) Add(index int, val int) {
	ov := findNP.nums2[index]
	findNP.nums2[index] += val
	findNP.countMap[ov]--
	findNP.countMap[findNP.nums2[index]]++
}

func (findNP *FindSumPairs) Count(tot int) int {
	cnt := 0
	for _, x := range findNP.nums1 {
		if v, has := findNP.countMap[tot-x]; has {
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
