package problems

const MAXNUM int = 1e5

var PRIMES1390 = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151,
	157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233,
	239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313,
}
var IS_PRIME1390 [MAXNUM + 1]bool

func initPirme1390() {
	if IS_PRIME1390[2] {
		return
	}
	for i := 2; i <= MAXNUM; i++ {
		IS_PRIME1390[i] = true
	}
	for _, p := range PRIMES1390 {
		start := p * p
		if start > MAXNUM {
			break
		}
		for j := start; j <= MAXNUM; j += p {
			IS_PRIME1390[j] = false
		}
	}
}

func sumDiv(x int, primes []int) int {
	if x < 4 {
		return 0
	}

	// Check if x = p^3
	for _, p := range primes {
		if p*p*p > x {
			break
		}
		if p*p*p == x {
			return 1 + p + p*p + p*p*p
		}
	}
	for _, p := range primes {
		if x%p == 0 {
			q := x / p
			if p == q || !IS_PRIME1390[q] {
				return 0
			}
			return 1 + x + p + q
		}
	}
	return 0
}

func sumFourDivisors(nums []int) int {
	initPirme1390()
	primes := PRIMES1390
	ans := 0
	for _, x := range nums {
		v := sumDiv(x, primes)
		ans += v
	}
	return ans
}
