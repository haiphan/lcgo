package problems

func minSubarray(nums []int, p int) int {
	n := len(nums)
	r := 0
	for _, x := range nums {
		r += x
		r %= p
	}
	if r == 0 {
		return 0
	}
	pm := make(map[int]int, n)
	ans := n
	cur := 0
	pm[0] = -1
	for i, x := range nums {
		if ans == 1 {
			return 1
		}
		cur = (cur + x) % p
		need := (cur - r + p) % p
		if l, has := pm[need]; has {
			ans = min(ans, i-l)
		}
		pm[cur] = i
	}
	if ans == n {
		return -1
	}
	return ans
}
