package problems

import (
	"testing"
)

func hasZero(n int) bool {
	if n == 0 {
		return true
	}
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}

func TestGetNoZeroIntegers(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "Example 1",
			n:    11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNoZeroIntegers(tt.n)
			a, b := got[0], got[1]
			if a+b != tt.n {
				t.Errorf("getNoZeroIntegers() sum = %v, want %v", a+b, tt.n)
			}
			if hasZero(a) || hasZero(b) {
				t.Errorf("getNoZeroIntegers() has zero in %v", got)
			}
		})
	}
}
