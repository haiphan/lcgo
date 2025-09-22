package problems

func maxFrequencyElements(nums []int) int {
	cm := make(map[int]int, len(nums))
	maxc := 0
	for _, x := range nums {
		cm[x]++
		maxc = max(maxc, cm[x])
	}
	res := 0
	for _, v := range cm {
		if v == maxc {
			res += maxc
		}
	}
	return res
}
