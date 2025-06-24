package problems

import (
	"testing"
)

func TestRangeSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		ranges [][]int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{-2, 0, 3, -5, 2, -1},
			ranges: [][]int{{0, 2}, {2, 5}, {0, 5}},
			want:   []int{1, -1, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := RangeSumConstructor(tt.nums)
			for i, r := range tt.ranges {
				got := rs.SumRange(r[0], r[1])
				if got != tt.want[i] {
					t.Errorf("SumRange(%d, %d) = %d, want %d", r[0], r[1], got, tt.want[i])
				}
			}
		})
	}
}
