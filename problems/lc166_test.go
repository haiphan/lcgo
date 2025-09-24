package problems

import (
	"reflect"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	tests := []struct {
		name        string
		numerator   int
		denominator int
		want        string
	}{
		{
			name:      "Example 1",
			numerator: 1, denominator: 2,
			want: "0.5",
		},
		{
			name:      "Example 2",
			numerator: 2, denominator: 1,
			want: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fractionToDecimal(tt.numerator, tt.denominator)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
