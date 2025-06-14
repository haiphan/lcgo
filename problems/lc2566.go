package problems

import "strconv"

func getMax(x int) int {
	s := strconv.Itoa(x)
	arr := []byte(s)
	c := byte('a')
	for i := range len(arr) {
		if c == 'a' && arr[i] != '9' {
			c = arr[i]
		}
		if c != 'a' && arr[i] == c {
			arr[i] = '9'
		}
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}

func getMin(x int) int {
	s := strconv.Itoa(x)
	arr := []byte(s)
	c := byte('a')
	for i := range len(arr) {
		if c == 'a' && arr[i] != '0' {
			c = arr[i]
		}
		if c != 'a' && arr[i] == c {
			arr[i] = '0'
		}
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}

func minMaxDifference(num int) int {
	return getMax(num) - getMin(num)
}
