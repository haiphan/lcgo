package problems

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	prev := ""
	res := make([]string, 0)
	for _, cur := range folder {
		if prev == "" || !strings.HasPrefix(cur, prev) {
			res = append(res, cur)
			prev = cur + "/"
		}
	}
	return res
}
