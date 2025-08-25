package problems

func separateDigits(nums []int) []int {
	N := len(nums)
	res := make([]int, 0, 3*N)
	addNum := func(x int) {
		if x == 0 {
			res = append(res, 0)
			return
		}
		for x > 0 {
			res = append(res, x%10)
			x /= 10
		}
	}
	for i := N - 1; i >= 0; i-- {
		addNum(nums[i])
	}
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
