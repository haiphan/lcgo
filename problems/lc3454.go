package problems

import "sort"

type YEvent struct {
	y      int
	x1, x2 int
	typ    int // 1 for bottom edge, -1 for top edge
}

// IntervalData stores the union width and cumulative area at each Y-step
type IntervalData struct {
	yStart  int
	width   int
	cumArea int64
}

// SegmentTree maintains the union length of active X-intervals
type XSegmentTree struct {
	count      []int
	widthTotal []int
	sortedX    []int
}

func (st *XSegmentTree) Update(node, start, end, l, r, val int) {
	// Standard Interval Segment Tree Update
	if l <= start && end <= r {
		st.count[node] += val
	} else {
		mid := (start + end) / 2
		if l < mid {
			st.Update(2*node, start, mid, l, r, val)
		}
		if r > mid {
			st.Update(2*node+1, mid, end, l, r, val)
		}
	}

	// Recalculate Width: If count > 0, the whole segment is covered
	if st.count[node] > 0 {
		st.widthTotal[node] = st.sortedX[end] - st.sortedX[start]
	} else if end-start > 1 {
		st.widthTotal[node] = st.widthTotal[2*node] + st.widthTotal[2*node+1]
	} else {
		st.widthTotal[node] = 0
	}
}

func separateSquares3454(squares [][]int) float64 {
	sqs := squares
	n := len(sqs)
	if n == 0 {
		return 0
	}

	// 1. Collect Events and unique X-coordinates
	events := make([]YEvent, 0, 2*n)
	xCoords := make([]int, 0, 2*n)
	for _, s := range sqs {
		x, y, l := s[0], s[1], s[2]
		events = append(events, YEvent{y, x, x + l, 1})
		events = append(events, YEvent{y + l, x, x + l, -1})
		xCoords = append(xCoords, x, x+l)
	}

	// 2. Efficient Coordinate Compression for X
	sort.Ints(xCoords)
	uniqueIdx := 0
	for i := 1; i < len(xCoords); i++ {
		if xCoords[i] != xCoords[uniqueIdx] {
			uniqueIdx++
			xCoords[uniqueIdx] = xCoords[i]
		}
	}
	sortedX := xCoords[:uniqueIdx+1]
	numX := len(sortedX)

	// Map X-values to their indices for O(1) lookup
	xMap := make(map[int]int, numX)
	for i, x := range sortedX {
		xMap[x] = i
	}

	// 3. Sort Events by Y
	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	// 4. Initialize Segment Tree
	tree := &XSegmentTree{
		count:      make([]int, 4*numX),
		widthTotal: make([]int, 4*numX),
		sortedX:    sortedX,
	}

	// 5. Sweep and Cache Cumulative Area
	history := make([]IntervalData, 0, len(events))
	var currentArea int64

	for i := 0; i < len(events)-1; i++ {
		tree.Update(1, 0, numX-1, xMap[events[i].x1], xMap[events[i].x2], events[i].typ)

		w := tree.widthTotal[1]
		dy := events[i+1].y - events[i].y

		if dy > 0 {
			currentArea += int64(w) * int64(dy)
			history = append(history, IntervalData{
				yStart:  events[i].y,
				width:   w,
				cumArea: currentArea,
			})
		}
	}

	if len(history) == 0 {
		return 0
	}

	// 6. Binary Search for the target interval
	totalArea := history[len(history)-1].cumArea
	target := float64(totalArea) / 2.0

	idx := sort.Search(len(history), func(i int) bool {
		return float64(history[i].cumArea) >= target
	})

	// 7. Solve for y0 using Linear Interpolation
	found := history[idx]
	areaBefore := float64(0)
	if idx > 0 {
		areaBefore = float64(history[idx-1].cumArea)
	}

	neededArea := target - areaBefore
	return float64(found.yStart) + (neededArea / float64(found.width))
}
