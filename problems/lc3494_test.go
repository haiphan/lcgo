package problems

import (
	"reflect"
	"testing"
)

func TestMinTime(t *testing.T) {
	tests := []struct {
		name  string
		skill []int
		mana  []int
		want  int64
	}{
		{
			name:  "Example 1",
			skill: []int{1, 5, 2, 4},
			mana:  []int{5, 1, 4, 2},
			want:  110,
		},
		{
			name:  "Example 2",
			skill: []int{1, 1, 1},
			mana:  []int{1, 1, 1},
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minTime(tt.skill, tt.mana)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
