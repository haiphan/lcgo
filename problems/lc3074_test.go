package problems

import (
	"reflect"
	"testing"
)

func TestMinimumBoxes(t *testing.T) {
	tests := []struct {
		name            string
		apple, capacity []int
		want            int
	}{
		{
			name:     "Example 1",
			apple:    []int{1, 3, 2},
			capacity: []int{4, 3, 1, 5, 2},
			want:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumBoxes(tt.apple, tt.capacity)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
