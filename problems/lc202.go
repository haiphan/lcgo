package problems

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	slow, fast := n, n
	for {
		slow, fast = transform(slow), transform(transform(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

func transform(n int) int {
	m := 0
	for n > 0 {
		d := n % 10
		m += d * d
		n /= 10
	}
	return m
}
