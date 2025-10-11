package problems

import (
	"reflect"
	"testing"
)

func TestMaximumTotalDamage(t *testing.T) {
	tests := []struct {
		name  string
		power []int
		want  int64
	}{
		{
			name:  "Example 1",
			power: []int{1, 1, 3, 4},
			want:  6,
		},
		{
			name:  "Example 2",
			power: []int{7, 1, 6, 6},
			want:  13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumTotalDamage(tt.power)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumTotalDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
