package problems

type Ride struct {
	e, t int
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	rs := make([][]Ride, n)
	for _, r := range rides {
		s, e, t := r[0], r[1], r[2]
		rs[s] = append(rs[s], Ride{e: e, t: e - s + t})
	}
	dp := make([]int, n+1)
	for i := n - 1; i >= 1; i-- {
		dp[i] = max(dp[i], dp[i+1])
		ra := rs[i]
		for _, r := range ra {
			dp[i] = max(dp[i], r.t+dp[r.e])
		}
	}
	return int64(dp[1])
}
