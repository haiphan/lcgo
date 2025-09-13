package problems

func toStrSet(arr []string) map[string]bool {
	res := make(map[string]bool, len(arr))
	for _, s := range arr {
		res[s] = true
	}
	return res
}
func findCommonResponse(responses [][]string) string {
	cm := make(map[string]int)
	res := ""
	maxF := 0
	for _, r := range responses {
		strSet := toStrSet(r)
		for k := range strSet {
			cm[k]++
			f := cm[k]
			if f > maxF || (f == maxF && k < res) {
				maxF = f
				res = k
			}
		}
	}
	return res
}
