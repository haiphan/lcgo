package problems

func subarrayBitwiseORs(arr []int) int {
	ns := make(map[int]bool)
	for i, x := range arr {
		ns[x] = true
		for j := i - 1; j >= 0 && (arr[j]|x) != arr[j]; j-- {
			arr[j] |= x
			ns[arr[j]] = true
		}
	}
	return len(ns)
}
