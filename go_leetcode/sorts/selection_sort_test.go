package sorts

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	listInts := [][]int{
		{1, 2, 3},
		{1, 6, 2, 0},
		{10, 9, 8, 7, 0, 18, 2, 2, 7, 0, 8},
	}

	for _, v := range listInts {
		selectionSort(v)
		fmt.Println(v)
	}
}
