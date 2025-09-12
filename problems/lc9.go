package problems

func isPalindrome(x int) bool {
	return x >= 0 && (x < 10 || x == rvInt(x))
}
