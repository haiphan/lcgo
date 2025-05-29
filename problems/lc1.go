package problems

// TwoSum finds two numbers in nums that add up to target.
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}