package problems

import (
	"reflect"
	"testing"
)

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		name   string
		fruits []int
		want   int
	}{
		{
			name:   "Example 1",
			fruits: []int{1, 2, 1},
			want:   3,
		},
		{
			name:   "Example 2",
			fruits: []int{0, 1, 2, 2},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalFruit(tt.fruits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
