package problems

func maxRemoval(nums []int, queries [][]int) int {
	N := len(nums)
	Q := len(queries)
	startCnt := make([]int, N+1)
	for _, q := range queries {
		startCnt[q[0]]++
	}
	for i := 1; i <= N; i++ {
		startCnt[i] += startCnt[i-1]
	}
	sortedStart := make([]int, Q)
	sortedEnd := make([]int, Q)
	for i := Q - 1; i >= 0; i-- {
		l, r := queries[i][0], queries[i][1]
		startCnt[l]--

		p := startCnt[l]
		sortedStart[p] = l
		sortedEnd[p] = r
	}
	availCnt := make([]int, N)
	runCnt := make([]int, N)
	avail := Q
	running := 0
	maxAvailEnd := -1
	minRunEnd := 0
	readPt := 0

	for p := 0; p < N; p++ {
		for readPt < Q && sortedStart[readPt] <= p {
			ep := sortedEnd[readPt]
			readPt++
			availCnt[ep]++
			if ep > maxAvailEnd {
				maxAvailEnd = ep
			}
		}
		for minRunEnd < p {
			cnt := runCnt[minRunEnd]
			if cnt != 0 {
				running -= cnt
				runCnt[minRunEnd] = 0
			}
			minRunEnd++
		}

		needed := nums[p] - running
		for needed > 0 {
			if maxAvailEnd < p {
				return -1
			}
			chosenEnd := maxAvailEnd
			availCnt[chosenEnd]--
			avail--
			if availCnt[chosenEnd] == 0 {
				for maxAvailEnd >= 0 && availCnt[maxAvailEnd] == 0 {
					maxAvailEnd--
				}
			}
			runCnt[chosenEnd]++
			running++
			needed--
		}
	}

	return avail
}
