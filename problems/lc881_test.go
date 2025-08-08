package problems

import (
	"reflect"
	"testing"
)

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name   string
		people []int
		limit  int
		want   int
	}{
		{
			name:   "Example 1",
			people: []int{1, 2},
			limit:  3,
			want:   1,
		},
		{
			name:   "Example 2",
			people: []int{3, 2, 2, 1},
			limit:  3,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numRescueBoats(tt.people, tt.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
