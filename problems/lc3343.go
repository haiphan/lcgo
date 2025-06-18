package problems

const MAXLEN = 80 // max length

var FACT_CACHE, IFACT_CACHE [MAXLEN + 1]int

func init() {
	FACT_CACHE[0] = 1
	for i := 1; i <= MAXLEN; i++ {
		FACT_CACHE[i] = FACT_CACHE[i-1] * i % MOD
	}
	IFACT_CACHE[MAXLEN] = powMod(FACT_CACHE[MAXLEN], MOD-2)
	for i := MAXLEN; i > 0; i-- {
		IFACT_CACHE[i-1] = IFACT_CACHE[i] * i % MOD
	}
}

func powMod(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 == 1 {
			res = res * x % MOD
		}
		x = x * x % MOD
	}
	return res
}

func countBalancedPermutations(num string) int {
	count := [10]int{}
	sum, n := 0, len(num)
	for _, c := range num {
		sum += int(c - '0')
	}
	if sum%2 == 1 {
		return 0
	}
	targetSum := sum / 2
	leftSize := n / 2
	NR, NC := targetSum+1, leftSize+1
	dp := make([][]int, NR)
	for i := range targetSum + 1 {
		dp[i] = make([]int, NC)
	}
	dp[0][0] = 1
	for _, c := range num {
		d := int(c - '0')
		count[d]++
		for i := targetSum; i >= d; i-- {
			prev := i - d
			for j := leftSize; j > 0; j-- {
				if dp[prev][j-1] > 0 {
					dp[i][j] = (dp[i][j] + dp[prev][j-1]) % MOD
				}
			}
		}
	}
	res := dp[targetSum][leftSize] * FACT_CACHE[leftSize] % MOD * FACT_CACHE[n-leftSize] % MOD
	for _, c := range count {
		if c > 0 {
			res = res * IFACT_CACHE[c] % MOD
		}
	}
	return res
}
