package problems

func minOperations2(nums1 []int, nums2 []int, k int) int64 {
	N := len(nums1)
	if k == 0 {
		for i := range N {
			if nums1[i] != nums2[i] {
				return -1
			}
		}
		return 0
	}
	inc, dec := 0, 0
	for i, x := range nums1 {
		d := x - nums2[i]
		isInc := d < 0
		if d < 0 {
			d = -d
		}
		if d%k != 0 {
			return -1
		}
		d /= k
		if isInc {
			inc += d
		} else {
			dec += d
		}
	}
	if dec == inc {
		return int64(inc)
	}
	return -1
}
