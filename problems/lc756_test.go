package problems

import (
	"reflect"
	"testing"
)

func TestPyramidTransition(t *testing.T) {
	tests := []struct {
		name    string
		bottom  string
		allowed []string
		want    bool
	}{
		{
			name:    "Example 1",
			bottom:  "BCD",
			allowed: []string{"BCC", "CDE", "CEA", "FFF"},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pyramidTransition(tt.bottom, tt.allowed)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pyramidTransition() = %v, want %v", got, tt.want)
			}
		})
	}
}
