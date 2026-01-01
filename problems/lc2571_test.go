package problems

import (
	"reflect"
	"testing"
)

func TestMinOperations2571(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    39,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations2571(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations2571() = %v, want %v", got, tt.want)
			}
		})
	}
}
