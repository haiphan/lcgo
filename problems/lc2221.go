package problems

var invTen [10]int = [10]int{0, 1, 0, 7, 0, 0, 0, 3, 0, 9}

func factor(x int) [3]int {
	exp2, exp5 := 0, 0
	for ; (x & 1) == 0; x = x >> 1 {
		exp2++
	}
	for ; (x % 5) == 0; x /= 5 {
		exp5++
	}
	x %= 10
	return [3]int{exp2, exp5, x}
}

func mulFactors(x, y [3]int) [3]int {
	return [3]int{x[0] + y[0], x[1] + y[1], x[2] * y[2] % 10}
}

func divFactors(x, y [3]int) [3]int {
	return [3]int{x[0] - y[0], x[1] - y[1], x[2] * invTen[y[2]] % 10}
}

func facTorsToInt(x [3]int) int {
	if x[0] > 0 && x[1] > 0 {
		return 0
	}
	if x[1] > 0 {
		return 5 * x[2] % 10
	}
	return (x[2] << x[0]) % 10
}

func Comb10(n int) []int {
	prev := [3]int{0, 0, 1}
	A := make([]int, n+1)
	A[0] = 1
	A[n] = 1
	for k := 1; k*2 <= n; k++ {
		cur := divFactors(mulFactors(prev, factor(n-k+1)), factor(k))
		A[k] = facTorsToInt(cur)
		A[n-k] = A[k]
		prev = cur
	}
	return A
}

func triangularSum(nums []int) int {
	n := len(nums) - 1
	A := Comb10(n)
	ans := 0
	for i := 0; i <= n; i++ {
		ans = (ans + A[i]*nums[i]) % 10
	}
	return ans
}
