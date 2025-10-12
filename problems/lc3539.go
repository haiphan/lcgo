package problems

import "math/bits"

var CNK [31][]int

func pascal() {
	if len(CNK[0]) == 1 {
		return
	}
	for i := range CNK {
		CNK[i] = make([]int, i+1)
		CNK[i][0] = 1
		half := i >> 1
		for j := 1; j <= i; j++ {
			if j <= half {
				CNK[i][j] = CNK[i-1][j-1] + CNK[i-1][j]
			} else {
				CNK[i][j] = CNK[i][i-j]
			}
		}
	}
}

func magicalSum(m int, k int, nums []int) int {
	pascal()
	L := len(nums)
	cc := make(map[int]int)
	var dfs func(remaining, odd_needed, index, carry int) int
	dfs = func(remaining, odd_needed, index, carry int) int {
		bc := bits.OnesCount32(uint32(carry))
		if remaining < 0 || odd_needed < 0 || (remaining+bc) < odd_needed {
			return 0
		}
		if remaining == 0 {
			if odd_needed == bc {
				return 1
			}
			return 0
		}
		if index >= L {
			return 0
		}
		ck := remaining + 31*odd_needed + 31*31*index + 31*31*51*carry
		cv, has := cc[ck]
		if has {
			return cv
		}
		ans := 0
		for take := 0; take <= remaining; take++ {
			ways := CNK[remaining][take] * intPowMod(nums[index], take, MOD) % MOD
			new_carry := carry + take
			ans = (ans + ways*dfs(remaining-take, odd_needed-(new_carry%2), index+1, new_carry>>1)) % MOD
		}
		cc[ck] = ans
		return ans
	}

	return dfs(m, k, 0, 0)
}
