package problems

func lengthOfLIS(nums []int) int {
	N := len(nums)
	seq := make([]int, 0, N)
	for _, x := range nums {
		if len(seq) == 0 || seq[len(seq)-1] < x {
			seq = append(seq, x)
		} else {
			for i := range seq {
				if seq[i] >= x {
					seq[i] = x
					break
				}
			}
		}
	}
	return len(seq)
}
