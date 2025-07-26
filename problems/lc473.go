package problems

import "sort"

func helper(start, edge, sum, target int, used []bool, arr []int) bool {
	if edge == 4 {
		return true
	}
	if sum > target {
		return false
	}
	if sum == target {
		return helper(0, edge+1, 0, target, used, arr)
	}
	for i := start; i < len(arr); i++ {
		if used[i] {
			continue
		}
		if i > start && !used[i-1] && arr[i-1] == arr[i] {
			continue
		}
		used[i] = true
		if helper(i+1, edge, sum+arr[i], target, used, arr) {
			return true
		}
		used[i] = false
	}
	return false
}

func makesquare(matchsticks []int) bool {
	mv, sum := 0, 0
	for _, x := range matchsticks {
		mv = max(mv, x)
		sum += x
	}
	if sum%4 != 0 {
		return false
	}
	target := sum / 4
	if target < mv {
		return false
	}
	sort.Ints(matchsticks)
	used := make([]bool, len(matchsticks))
	return helper(0, 0, 0, target, used, matchsticks)
}
