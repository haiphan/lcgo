package problems

func longestSubarray(nums []int) int {
	N := len(nums)
	i := 0
	maxv := nums[0]
	for _, x := range nums {
		maxv = max(maxv, x)
	}
	res := 1
	for i < N {
		for i < N && nums[i] < maxv {
			i++
		}
		if i == N {
			break
		}
		j := i + 1
		for j < N && nums[j] == maxv {
			j++
		}
		res = max(res, j-i)
		i = j + 1
	}
	return res
}
