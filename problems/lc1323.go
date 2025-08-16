package problems

func maximum69Number(num int) int {
	p10s := [4]int{1, 10, 100, 1000}
	for i := 3; i >= 0; i-- {
		p10 := p10s[i]
		if num < p10 {
			continue
		}
		d := (num / p10) % 10
		if d < 9 {
			return num + (9-d)*p10
		}
	}
	return num
}
