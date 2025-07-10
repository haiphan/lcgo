package problems

import (
	"reflect"
	"testing"
)

func TestMaxFreeTime2(t *testing.T) {
	tests := []struct {
		name      string
		eventTime int
		startTime []int
		endTime   []int
		want      int
	}{
		{
			name:      "Example 1",
			eventTime: 5,
			startTime: []int{1, 3},
			endTime:   []int{2, 5},
			want:      2,
		},
		{
			name:      "Example 2",
			eventTime: 10,
			startTime: []int{0, 7, 9},
			endTime:   []int{1, 8, 10},
			want:      7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFreeTime2(tt.eventTime, tt.startTime, tt.endTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxFreeTime2() = %v, want %v", got, tt.want)
			}
		})
	}
}
