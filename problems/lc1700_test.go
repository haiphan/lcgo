package problems

import (
	"reflect"
	"testing"
)

func TestCountStudents(t *testing.T) {
	tests := []struct {
		name                 string
		students, sandwiches []int
		want                 int
	}{
		{
			name:       "Example 1",
			students:   []int{1, 1, 0, 0},
			sandwiches: []int{0, 1, 0, 1},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countStudents(tt.students, tt.sandwiches)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
