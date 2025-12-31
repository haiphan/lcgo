package problems

func countStudents(students []int, sandwiches []int) int {
	cnt := []int{0, 0}
	for _, s := range students {
		cnt[s]++
	}
	for _, s := range sandwiches {
		// If no student wants this sandwich, break
		if cnt[s] == 0 {
			break
		}
		// A student takes this sandwich
		cnt[s]--
	}
	return cnt[0] + cnt[1]
}
