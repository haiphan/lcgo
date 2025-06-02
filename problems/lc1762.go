package problems

func tupleSameProduct(nums []int) int {
	N := len(nums)
	pm := make(map[int]int, N*N/4)
	for i := range N {
		for j := i + 1; j < N; j++ {
			pm[nums[i]*nums[j]]++
		}
	}
	res := 0
	for _, v := range pm {
		res += v * (v - 1)
	}
	return res * 4
}
