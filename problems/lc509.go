package problems

func fib(n int) int {
	if n <= 1 {
		return n
	}
	//p2,p1,cur
	p2, p1 := 0, 1
	for i := 2; i <= n; i++ {
		cur := p1 + p2
		p1, p2 = cur, p1
	}
	return p1
}
