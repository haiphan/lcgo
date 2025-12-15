package problems

import (
	"reflect"
	"testing"
)

func TestClumsy(t *testing.T) {
	tests := []struct {
		name    string
		n, want int
	}{
		{
			name: "Example 1",
			n:    4,
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clumsy(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}
