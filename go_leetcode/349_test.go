package go_leetcode

import "testing"

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		list []int
		k    int
		want int
	}{
		{[]int{3, 1, 6, 2, 8, 4, 5}, 3, 5},
		{[]int{-5, -3, -9, -4, -7}, 1, -3},
		{[]int{}, 5, -1},
		{[]int{6, 5, 4, 3, 2, 1, 6}, 2, 6},
	}

	for _, tc := range testCases {
		got := findKthLargest(tc.list, tc.k)
		if got != tc.want {
			t.Errorf("findKthLargest(%v, %v) = %v; want %v", tc.list, tc.k, got, tc.want)
		}
	}
}
