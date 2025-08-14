package problems

func sortPermutation(nums []int) int {
	BIG := 0x1ffff
	km := BIG
	for i, x := range nums {
		if i != x {
			km &= x
		}
	}
	if km == BIG {
		return 0
	}
	return km
}
