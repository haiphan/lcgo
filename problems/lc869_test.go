package problems

import (
	"reflect"
	"testing"
)

func TestReorderedPowerOf2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    1,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reorderedPowerOf2(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
