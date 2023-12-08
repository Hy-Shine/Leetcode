package go_leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{
			name: "only",
			list: []int{0, 1, 2},
			want: []int{0, 1, 4},
		},
		{
			name: "both",
			list: []int{-3, -2, -1, 0, 1, 2},
			want: []int{0, 1, 1, 4, 4, 9},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sortedSquares(test.list)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("sortedSquares(%v) = %v; but want %v", test.list, got, test.want)
			}
		})
	}
}
