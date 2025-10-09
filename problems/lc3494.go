package problems

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	prefix_skills := make([]int, n)
	for i := 1; i < n; i++ {
		prefix_skills[i] = prefix_skills[i-1] + skill[i]
	}
	total_time := skill[0] * mana[0]

	for j := 1; j < m; j++ {
		max_time := skill[0] * mana[j]
		for i := 1; i < n; i++ {
			diff_time := prefix_skills[i]*mana[j-1] - prefix_skills[i-1]*mana[j]
			max_time = max(max_time, diff_time)
		}
		total_time += max_time
	}

	return int64(total_time + prefix_skills[n-1]*mana[m-1])
}
