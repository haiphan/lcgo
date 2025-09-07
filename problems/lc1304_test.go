package problems

import (
	"testing"
)

func TestSumZero(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "Example 1",
			n:    5,
		},
		{
			name: "Example 6",
			n:    12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumZero(tt.n)
			sum := 0
			ns := make(map[int]bool, 0)
			for _, v := range got {
				ns[v] = true
				sum += v
			}
			if sum != 0 {
				t.Errorf("SumZero() sum = %v, want 0", sum)
			}
			if len(ns) != tt.n {
				t.Errorf("SumZero() len = %v, want %v", len(ns), tt.n)
			}
		})
	}
}
