package problems

func maxProfit3652(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	h := k >> 1
	ss := strategy
	pp := prices
	for i := range ss {
		ss[i] *= prices[i]
	}
	sum := 0
	hs := 0
	win_orig := 0
	win_mod := 0
	for i, x := range ss {
		if i >= h {
			hs += pp[i]
		}
		if i < k {
			win_orig += x
		}
		if i >= h && i < k {
			win_mod += pp[i]
		}
		sum += x
	}

	if n == k {
		change := hs - sum
		return int64(sum + max(0, change))
	}
	max_ch := win_mod - win_orig

	for i := 1; i <= n-k; i++ {
		win_orig += ss[i+k-1] - ss[i-1]
		win_mod += pp[i+k-1] - pp[i-1+h]
		max_ch = max(max_ch, win_mod-win_orig)
	}
	return int64(sum + max(0, max_ch))
}
