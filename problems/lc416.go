package problems

import "math/big"

func canPartition(nums []int) bool {
	sum := 0
	maxN := 0
	for _, n := range nums {
		sum += n
		maxN = max(maxN, n)
	}
	if sum%2 == 1 {
		return false
	}
	half := sum / 2
	if maxN >= half {
		return maxN == half
	}
	dp := new(big.Int)
	dp.Lsh(big.NewInt(1), uint(half))
	for _, n := range nums {
		temp := new(big.Int)
		dp.Or(dp, temp.Rsh(dp, uint(n)))
	}
	return dp.Bit(0) == 1
}
