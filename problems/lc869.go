package problems

func digitHash(n int) int {
	hash := 0
	for n > 0 {
		d := n % 10
		hash += 1 << (d * 3)
		n /= 10
	}
	return hash
}
func reorderedPowerOf2(n int) bool {
	sn := digitHash(n)
	for i := range 31 {
		if sn == digitHash(1<<i) {
			return true
		}
	}
	return false
}
