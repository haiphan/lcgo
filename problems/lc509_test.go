package problems

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
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
			n:    3,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fib(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
