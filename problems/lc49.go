package problems

var BASE int64 = 29

func getHash(s string) int64 {
	charCounts := [26]int64{}
	for _, c := range s {
		charCounts[c-'a']++
	}
	var h int64 = 0
	for _, cnt := range charCounts {
		h = BASE*h + (cnt + 1)
	}
	return h
}

func groupAnagrams(strs []string) [][]string {
	anaMap := make(map[int64][]string)
	for _, s := range strs {
		h := getHash(s)
		anaMap[h] = append(anaMap[h], s)
	}
	res := [][]string{}
	for _, value := range anaMap {
		res = append(res, value)
	}
	return res
}
