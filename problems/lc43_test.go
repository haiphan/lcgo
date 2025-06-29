package problems

import (
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{
			name: "Example 1",
			num1: "123",
			num2: "456",
			want: "56088",
		},
		{
			name: "Example 2",
			num1: "2",
			num2: "3",
			want: "6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := multiply(tt.num1, tt.num2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
