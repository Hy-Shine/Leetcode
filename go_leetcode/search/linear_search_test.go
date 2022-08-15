package search

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	listInts := [][]int{
		{1, 2, 3, 7, 9, 10, 20},
		{10, 19, 20, 78, 65, 26},
	}

	target := 20
	for _, v := range listInts {
		index := LinearSearch(v, target)
		fmt.Printf("target %d index is %d\n", target, index)
	}
}
