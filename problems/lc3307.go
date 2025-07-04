package problems

func kthCharacter2(k int64, operations []int) byte {
	k--
	i, cnt := 0, 0
	for k > 0 {
		cnt += (int(k) & 1 & operations[i])
		k = k >> 1
		i++
	}
	return 'a' + byte(cnt%26)
}
