package problems

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	prefix_skills := make([]int, n)
	prefix_skills[0] = skill[0]
	for i := 1; i < n; i++ {
		prefix_skills[i] = prefix_skills[i-1] + skill[i]
	}
	total_time := 0

	for j := 1; j < m; j++ {
		max_time := 0
		cur, prev := mana[j], mana[j-1]
		e := 0
		for i := range n {
			max_time = max(max_time, prefix_skills[i]*prev-e)
			e = prefix_skills[i] * cur
		}
		total_time += max_time
	}

	return int64(total_time + prefix_skills[n-1]*mana[m-1])
}
