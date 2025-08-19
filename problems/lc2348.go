package problems

func zeroFilledSubarray(nums []int) int64 {
	res, cnt := 0, 0
	for _, x := range nums {
		if x == 0 {
			cnt++
		} else {
			cnt = 0
		}
		res += cnt
	}
	return int64(res)
}
