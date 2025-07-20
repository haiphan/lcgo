package problems

import (
	"reflect"
	"testing"
)

func TestMaxProfitAssignment(t *testing.T) {
	tests := []struct {
		name       string
		difficulty []int
		profit     []int
		worker     []int
		want       int
	}{
		{
			name:       "Example 1",
			difficulty: []int{2, 4, 6, 8, 10},
			profit:     []int{10, 20, 30, 40, 50},
			worker:     []int{4, 5, 6, 7},
			want:       100,
		},
		{
			name:       "Example 2",
			difficulty: []int{85, 47, 57},
			profit:     []int{24, 66, 99},
			worker:     []int{40, 25, 25},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfitAssignment(tt.difficulty, tt.profit, tt.worker)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxProfitAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
