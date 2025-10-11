package problems

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	pre_skill := make([]int, n+1)
	for i, x := range skill {
		pre_skill[i+1] = pre_skill[i] + x
	}
	dec := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		if skill[i] > skill[dec[len(dec)-1]] {
			dec = append(dec, i)
		}
	}
	inc := []int{0}
	for i := 1; i < n; i++ {
		if skill[i] > skill[inc[len(inc)-1]] {
			inc = append(inc, i)
		}
	}
	start := 0
	for j := 1; j < m; j++ {
		stack := dec
		if mana[j-1] < mana[j] {
			stack = inc
		}
		maxv := 0
		for _, i := range stack {
			maxv = max(maxv, mana[j-1]*pre_skill[i+1]-mana[j]*pre_skill[i])
		}
		start += maxv
	}
	return int64(start + pre_skill[n]*mana[m-1])
}
