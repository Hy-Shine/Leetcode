package go_leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result []int
	}{
		{
			name:   "No zeros",
			nums:   []int{1, 2, 3, 4},
			result: []int{1, 2, 3, 4},
		},
		{
			name:   "All zeros",
			nums:   []int{0, 0, 0, 0},
			result: []int{0, 0, 0, 0},
		},
		{
			name:   "Mixed zeros and non-zeros",
			nums:   []int{1, 0, 2, 0, 3, 0, 4},
			result: []int{1, 2, 3, 4, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			moveZeroes(test.nums)
			if !reflect.DeepEqual(test.nums, test.result) {
				t.Errorf("Expected %v, got %v", test.result, test.nums)
			}
		})
	}
}
