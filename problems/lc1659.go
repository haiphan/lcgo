package problems

func maximumUniqueSubarray(nums []int) int {
	N := len(nums)
	var seen [10001]bool
	var sum, res, l int
	for r := range N {
		numR := nums[r]
		sum += numR
		if seen[numR] {
			for l <= r {
				numL := nums[l]
				seen[numL] = false
				sum -= numL
				l++
				if numL == numR {
					break
				}
			}
		}
		seen[numR] = true
		res = max(res, sum)
	}
	return res
}
