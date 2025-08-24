package problems

func minimumIndex(nums []int) int {
	n := len(nums)
	dom := majorityElement(nums)
	dc := 0
	for _, v := range nums {
		if v == dom {
			dc++
		}
	}
	l := 0
	for i, v := range nums {
		if v == dom {
			l++
		}
		if l > ((i+1)>>1) && (dc-l) > ((n-i-1)>>1) {
			return i
		}
	}
	return -1
}
