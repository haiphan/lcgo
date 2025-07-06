package problems

import (
	"testing"
)

type SumPair struct {
	op    string
	index int
	val   int
	tot   int
	want  int
}

func TestFindSumPairs(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		commands []SumPair
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 1, 2, 2, 2, 3},
			nums2: []int{1, 4, 5, 2, 5, 4},
			commands: []SumPair{
				{op: "count", tot: 7, want: 8},
				{op: "add", index: 3, val: 2},
				{op: "count", tot: 8, want: 2},
				{op: "count", tot: 4, want: 1},
				{op: "add", index: 0, val: 1},
				{op: "add", index: 1, val: 1},
				{op: "count", tot: 7, want: 11},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := FindSumPairsConstructor(tt.nums1, tt.nums2)
			for _, cmd := range tt.commands {
				switch cmd.op {
				case "count":
					got := sp.Count(cmd.tot)
					if got != cmd.want {
						t.Errorf("Count(%d) = %d, want %d", cmd.tot, got, cmd.want)
					}
				case "add":
					sp.Add(cmd.index, cmd.val)
				default:
					t.Errorf("Unknown operation: %s", cmd.op)
				}
			}
		})
	}
}
