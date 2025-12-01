package problems

import (
	"reflect"
	"testing"
)

func TestMaxRunTime(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		batteries []int
		want      int64
	}{
		{
			name:      "Example 1",
			n:         2,
			batteries: []int{3, 3, 3},
			want:      4,
		},
		{
			name:      "Example 2",
			n:         2,
			batteries: []int{1, 1, 1, 1},
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxRunTime(tt.n, tt.batteries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxRunTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
