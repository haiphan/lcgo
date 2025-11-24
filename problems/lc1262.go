package problems

func maxSumDivThree(nums []int) int {
	const maxn int = 1e4
	n := len(nums)
	big := maxn*n + 1
	sum := 0
	min1 := big
	min2 := big
	for _, x := range nums {
		sum += x
		r := x % 3
		switch r {
		case 1:
			min1, min2 = min(min1, x), min(min2, min1+x)
		case 2:
			min1, min2 = min(min1, min2+x), min(min2, x)
		}
	}
	r := sum % 3
	if r == 0 {
		return sum
	}
	if r == 1 {
		return sum - min1
	}
	return sum - min2
}
