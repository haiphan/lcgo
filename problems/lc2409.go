package problems

import (
	"strconv"
	"strings"
)

var MDAYS = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func dayToInt(s string) int {
	mmdd := strings.Split(s, "-")
	m, _ := strconv.Atoi(mmdd[0])
	d, _ := strconv.Atoi(mmdd[1])
	for i := 1; i < m; i++ {
		d += MDAYS[i-1]
	}
	return d
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	s1, e1 := dayToInt(arriveAlice), dayToInt(leaveAlice)
	s2, e2 := dayToInt(arriveBob), dayToInt(leaveBob)
	s := max(s1, s2)
	e := min(e1, e2)
	return max(0, e-s+1)
}
