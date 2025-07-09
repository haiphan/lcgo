package problems

import (
	"reflect"
	"testing"
)

func TestMaxFreeTime(t *testing.T) {
	tests := []struct {
		name      string
		eventTime int
		k         int
		startTime []int
		endTime   []int
		want      int
	}{
		{
			name:      "Example 1",
			eventTime: 5,
			k:         1,
			startTime: []int{1, 3},
			endTime:   []int{2, 5},
			want:      2,
		},
		{
			name:      "Example 2",
			eventTime: 10,
			k:         1,
			startTime: []int{0, 2, 9},
			endTime:   []int{1, 4, 10},
			want:      6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFreeTime(tt.eventTime, tt.k, tt.startTime, tt.endTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
