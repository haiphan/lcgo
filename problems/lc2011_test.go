package problems

import (
	"reflect"
	"testing"
)

func TestFinalValueAfterOperations(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		want       int
	}{
		{
			name:       "Example 1",
			operations: []string{"--X", "X++", "X++"},
			want:       1,
		},
		{
			name:       "Example 2",
			operations: []string{"++X", "++X", "X++"},
			want:       3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finalValueAfterOperations(tt.operations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
