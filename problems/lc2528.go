package problems

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	dd := make([]int, n+1)
	dk := make([]int, n+1)
	tr := r << 1
	left := stations[0]
	right := k
	for i, v := range stations {
		left = min(left, v)
		right += v
		dd[max(0, i-r)] += v
		dd[min(n, i+r+1)] -= v
	}
	canReach := func(lb int) bool {
		copy(dk, dd)
		cur, remain := 0, k
		for i := range stations {
			cur += dk[i]
			if cur < lb {
				need := lb - cur
				if need > remain {
					return false
				}
				remain -= need
				cur = lb
				dk[min(n, i+tr+1)] -= need
			}
		}

		return true
	}

	ans := left
	for left <= right {
		mid := (left + right) >> 1
		if canReach(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int64(ans)
}
