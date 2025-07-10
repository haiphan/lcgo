package problems

func maxFreeTime2(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	canMove := make([]bool, n)
	maxFree := 0
	maxL, maxR := 0, 0
	getDur := func(i int) int {
		return endTime[i] - startTime[i]
	}
	for i := range n {
		ri := n - 1 - i
		if !canMove[i] && getDur(i) <= maxL {
			canMove[i] = true
		}
		if i == 0 {
			maxL = startTime[0]
		} else {
			maxL = max(maxL, startTime[i]-endTime[i-1])
		}
		if !canMove[ri] && getDur(ri) <= maxR {
			canMove[ri] = true
		}
		if ri == n-1 {
			maxR = eventTime - endTime[ri]
		} else {
			maxR = max(maxR, startTime[ri+1]-endTime[ri])
		}
	}
	for i := range n {
		l := 0
		if i > 0 {
			l = endTime[i-1]
		}
		r := eventTime
		if i != (n - 1) {
			r = startTime[i+1]
		}
		cand := r - l
		if !canMove[i] {
			cand -= getDur(i)
		}
		maxFree = max(maxFree, cand)
	}
	return maxFree
}
