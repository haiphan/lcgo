package problems

func specialTriplets(nums []int) int {
	const BIG int = 100001
	var ci [BIG]int
	var ck [BIG]int
	ans := 0
	for _, x := range nums {
		ans += ck[x]
		y := x << 1
		if y < BIG {
			ck[y] += ci[y]
		}
		ci[x]++
	}
	ans %= MOD
	return ans
}
