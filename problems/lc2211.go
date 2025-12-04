package problems

func countCollisions(directions string) int {
	dd := directions
	n := len(dd)
	l := 0
	for l < n && dd[l] == 'L' {
		l++
	}
	// ans := 0
	r := n - 1
	for r >= l && dd[r] == 'R' {
		r--
	}
	cs := 0
	for i := l; i <= r; i++ {
		if dd[i] == 'S' {
			cs++
		}
	}
	return r - l + 1 - cs
}
