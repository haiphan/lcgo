package problems

import (
	"reflect"
	"testing"
)

func TestCountBalancedPermutations(t *testing.T) {
	tests := []struct {
		name string
		num  string
		want int
	}{
		{
			name: "Example 1",
			num:  "123",
			want: 2,
		},
		{
			name: "Example 2",
			num:  "112",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBalancedPermutations(tt.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBalancedPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
