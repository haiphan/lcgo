package problems

func maxDistance624(arrays [][]int) int {
	n := len(arrays)
	ans := 0
	minv := arrays[0][0]
	maxv := arrays[0][len(arrays[0])-1]
	for i := 1; i < n; i++ {
		arr := arrays[i]
		ans = max(ans, abs(arr[len(arr)-1]-minv), abs(maxv-arr[0]))
		minv = min(minv, arr[0])
		maxv = max(maxv, arr[len(arr)-1])
	}
	return ans
}
