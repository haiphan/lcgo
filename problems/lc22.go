package problems

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	L := 2 * n
	var bt func(c string, open, close int)
	bt = func(c string, open, close int) {
		if len(c) == L {
			res = append(res, c)
			return
		}
		if open < n {
			bt(c+"(", open+1, close)
		}
		if open > close {
			bt(c+")", open, close+1)
		}
	}
	bt("", 0, 0)
	return res
}
