package go_leetcode

import (
	"testing"
)

func TestMaxArrayValue(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int64
	}{
		{"Descending Order", []int{5, 4, 3, 2, 1}, 5},
		{"Ascending Order", []int{1, 2, 3, 4, 5}, 15},
		{"one element", []int{1}, 1},
		{"two elements", []int{1, 2}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArrayValue(tt.input)
			if got != tt.want {
				t.Errorf("maxArrayValue(%v) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
