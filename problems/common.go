package problems

import "lcgo/utils"

const MOD int = 1e9 + 7

type TreeNode = utils.TreeNode

func intPow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	res := pow(x*x, n/2)
	if n%2 == 1 {
		res *= x
	}
	return res
}
