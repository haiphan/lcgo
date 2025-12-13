package problems

import (
	"reflect"
	"testing"
)

func TestValidateCoupons(t *testing.T) {
	tests := []struct {
		name               string
		code, businessLine []string
		isActive           []bool
		want               []string
	}{
		{
			name:         "Example 1",
			code:         []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
			businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
			isActive:     []bool{true, true, true, true},
			want:         []string{"PHARMA5", "SAVE20"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateCoupons(tt.code, tt.businessLine, tt.isActive)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}
