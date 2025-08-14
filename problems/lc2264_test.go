package problems

import (
	"reflect"
	"testing"
)

func TestLargestGoodInteger(t *testing.T) {
	tests := []struct {
		name string
		num  string
		want string
	}{
		{
			name: "Example 1",
			num:  "6777133339",
			want: "777",
		},
		{
			name: "Example 2",
			num:  "2300019",
			want: "000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestGoodInteger(tt.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
