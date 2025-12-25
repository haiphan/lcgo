package problems

import "sort"

var EVCACHE, EVBUF [200000]int

func radixSortMSBOnly(arr []int) {
	n := len(arr)
	buffer := EVBUF[:n]

	const bitsPerPass = 11
	const bucketSize = 1 << bitsPerPass
	const mask = bucketSize - 1
	swap := false
	// Start at bit 20, end at bit 51 (31 bits total)
	for shift := uint(20); shift < 51; shift += bitsPerPass {
		var count [bucketSize]int

		for _, v := range arr {
			count[(v>>shift)&mask]++
		}

		pos := 0
		for i := 0; i < bucketSize; i++ {
			tmp := count[i]
			count[i] = pos
			pos += tmp
		}

		for _, v := range arr {
			idx := (v >> shift) & mask
			buffer[count[idx]] = v
			count[idx]++
		}

		arr, buffer = buffer, arr
		swap = !swap
	}
	if swap {
		copy(buffer, arr)
	}
}

func customSort(arr []int) {
	if len(arr) < 1<<9 {
		sort.Ints(arr)
	} else {
		radixSortMSBOnly(arr)
	}
}

func maxTwoEvents(events [][]int) int {
	n := len(events)
	m := n << 1
	arr := EVCACHE[:m]
	eMask := 1 << 20
	vMask := eMask - 1
	for i := range events {
		s, e, v := events[i][0], events[i][1], events[i][2]
		arr[i*2] = s<<21 | v
		arr[i*2+1] = e<<21 | eMask | v
	}
	customSort(arr)
	ans, maxv := 0, 0
	for _, t := range arr {
		v := t & vMask
		if (t & eMask) == 0 {
			ans = max(ans, maxv+v)
		} else {
			maxv = max(maxv, v)
		}
	}
	return ans
}
