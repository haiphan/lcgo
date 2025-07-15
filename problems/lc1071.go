package problems

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:GCD(len(str1), len(str2))]
}
