package problems

import (
	"sort"
	"unicode"
)

func isValidBL(s string) bool {
	return s == "electronics" || s == "grocery" || s == "pharmacy" || s == "restaurant"
}
func isValidCode(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if r != '_' && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func blToRank(s string) int {
	switch s {
	case "electronics":
		return 0
	case "grocery":
		return 1
	case "pharmacy":
		return 2
	default:
		return 3
	}
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	n := len(code)
	arr := make([][2]string, 0, n)
	for i, c := range code {
		if isActive[i] && isValidCode(c) && isValidBL(businessLine[i]) {
			arr = append(arr, [2]string{c, businessLine[i]})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		l1, l2 := blToRank(arr[i][1]), blToRank(arr[j][1])
		if l1 == l2 {
			return arr[i][0] < arr[j][0]
		}
		return l1 < l2
	})
	ans := make([]string, len(arr))
	for i := range arr {
		ans[i] = arr[i][0]
	}
	return ans
}
