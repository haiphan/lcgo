package problems

import (
	"reflect"
	"testing"
)

func TestIntegerBreak(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    2,
			want: 1,
		},
		{
			name: "Example 2",
			n:    10,
			want: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := integerBreak(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
