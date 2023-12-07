package go_leetcode

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		list   []int
		target int
		want   int
	}{
		{
			name:   "Empty List",
			list:   []int{},
			target: 5,
			want:   -1,
		},
		{
			name:   "Target in List",
			list:   []int{1, 3, 5, 7, 9},
			target: 5,
			want:   2,
		},
		{
			name:   "Target not in List",
			list:   []int{1, 3, 5, 7, 9},
			target: 6,
			want:   -1,
		},
		{
			name:   "Target at Start",
			list:   []int{1, 3, 5, 7, 9},
			target: 1,
			want:   0,
		},
		{
			name:   "Target at End",
			list:   []int{1, 3, 5, 7, 9},
			target: 9,
			want:   4,
		},
		{
			name:   "Single Element List",
			list:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "Target less than all elements",
			list:   []int{5, 10, 15, 20},
			target: 3,
			want:   -1,
		},
		{
			name:   "Target greater than all elements",
			list:   []int{5, 10, 15, 20},
			target: 25,
			want:   -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := binarySearch(test.list, test.target)
			if got != test.want {
				t.Errorf("binarySearch(%v, %d) = %d; want %d", test.list, test.target, got, test.want)
			}
		})
	}
}
