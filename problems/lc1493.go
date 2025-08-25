package problems

func longestSubarrayOfOne(nums []int) int {
	prev, cnt, res := 0, 0, 0
	hasz := false
	for _, x := range nums {
		if x == 0 {
			hasz = true
			res = max(res, prev+cnt)
			prev = cnt
			cnt = 0
		} else {
			cnt++
		}
	}
	res = max(res, prev+cnt)
	if hasz {
		return res
	}
	return res - 1
}
