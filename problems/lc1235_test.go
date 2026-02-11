package problems

import (
	"reflect"
	"testing"
)

func TestJobScheduling(t *testing.T) {
	tests := []struct {
		name                       string
		startTime, endTime, profit []int
		want                       int
	}{
		{
			name:      "Example 1",
			startTime: []int{1, 2, 3, 3},
			endTime:   []int{3, 4, 5, 65},
			profit:    []int{50, 10, 40, 70},
			want:      120,
		},
		{
			name:      "Example 2",
			startTime: []int{1, 2, 3, 4, 6},
			endTime:   []int{3, 5, 10, 6, 9},
			profit:    []int{20, 20, 100, 70, 60},
			want:      150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jobScheduling(tt.startTime, tt.endTime, tt.profit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jobScheduling() = %v, want %v", got, tt.want)
			}
		})
	}
}
