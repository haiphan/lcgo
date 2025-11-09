package problems

import (
	"reflect"
	"testing"
)

func TestCountOperations(t *testing.T) {
	tests := []struct {
		name       string
		num1, num2 int
		want       int
	}{
		{
			name: "Example 1",
			num1: 2,
			num2: 3,
			want: 3,
		},
		{
			name: "Example 2",
			num1: 10,
			num2: 10,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countOperations(tt.num1, tt.num2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
