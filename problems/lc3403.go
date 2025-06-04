package problems

func answerString(word string, numFriends int) string {
	L := len(word)
	if numFriends == 1 {
		return word
	}
	ML := L - numFriends + 1
	res := ""
	for i := range L {
		end := min(L, i+ML)
		sub := word[i:end]
		if len(res) == 0 || res < sub {
			res = sub
		}
	}
	return res
}
