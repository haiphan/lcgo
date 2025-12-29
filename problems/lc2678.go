package problems

import "strconv"

func countSeniors(details []string) int {
	// 0-9 phone, 10 g, 11-12 age
	ans := 0
	for _, d := range details {
		v, _ := strconv.Atoi(d[11:13])
		if v > 60 {
			ans++
		}
	}
	return ans
}
