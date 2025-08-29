package problems

func maximumProduct(nums []int) int {
	// min1,min2...,max1,max2,max3
	min1, min2 := 1001, 1001
	max1, max2, max3 := -1001, -1001, -1001
	for _, x := range nums {
		if x < min1 {
			min1, min2 = x, min1
		} else if x < min2 {
			min2 = x
		}
		if x > max3 {
			max1, max2, max3 = max2, max3, x
		} else if x > max2 {
			max1, max2 = max2, x
		} else if x > max1 {
			max1 = x
		}
	}

	return max(min1*min2*max3, max1*max2*max3)
}
