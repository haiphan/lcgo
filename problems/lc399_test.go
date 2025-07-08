package problems

import (
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			name: "Example 1",
			equations: [][]string{
				{"a", "b"},
				{"b", "c"},
			},
			values: []float64{2.0, 3.0},
			queries: [][]string{
				{"a", "c"},
				{"b", "a"},
				{"a", "e"},
				{"a", "a"},
				{"x", "x"},
			},
			want: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name: "Example 2",
			equations: [][]string{
				{"a", "b"},
				{"b", "c"},
				{"bc", "cd"},
			},
			values: []float64{1.5, 2.5, 5.0},
			queries: [][]string{
				{"a", "c"},
				{"c", "b"},
				{"bc", "cd"},
				{"cd", "bc"},
			},
			want: []float64{3.75, 0.4, 5.0, 0.2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcEquation(tt.equations, tt.values, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
