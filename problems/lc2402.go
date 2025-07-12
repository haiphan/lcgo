package problems

import "sort"

const MAXROOM int = 101

func getRoom(x int) int {
	return x % MAXROOM
}

func getEnd(x int) int {
	return x / MAXROOM
}

func mostBooked(n int, meetings [][]int) int {
	L := len(meetings)
	cm := make([]int, n)
	rq := make([]int, n)
	mq := make([]int, 0, L)
	maxCnt := 0
	for i := range n {
		rq[i] = i
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	for _, m := range meetings {
		start, end := m[0], m[1]
		dur := end - start
		for len(mq) > 0 && getEnd(mq[0]) <= start {
			room := getRoom(mq[0])
			mq = hPop(mq)
			rq = hPush(rq, room)
		}
		if len(rq) == 0 {
			dm := mq[0]
			mq = hPop(mq)
			rq = hPush(rq, getRoom(dm))
			end = getEnd(dm) + dur
		}
		room := rq[0]
		rq = hPop(rq)
		mq = hPush(mq, MAXROOM*end+room)
		cm[room]++
		maxCnt = max(maxCnt, cm[room])
	}
	for i := range n {
		if cm[i] == maxCnt {
			return i
		}
	}
	return 0
}
