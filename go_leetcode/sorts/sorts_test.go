package sorts

import (
	"reflect"
	"testing"
)

// 选择排序

func TestSelectionSort(t *testing.T) {
	t.Run("sorting already sorted array", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := selectionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("sorting array in descending order", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		result := selectionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("sorting array with negative numbers", func(t *testing.T) {
		input := []int{3, -1, 5, -2, 0}
		expected := []int{-2, -1, 0, 3, 5}
		result := selectionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("sorting array with duplicate elements", func(t *testing.T) {
		input := []int{3, 5, 2, 3, 2}
		expected := []int{2, 2, 3, 3, 5}
		result := selectionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "Already sorted list",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Reverse sorted list",
			input:  []int{5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "List with duplicate elements",
			input:  []int{3, 5, 2, 5, 1, 3},
			output: []int{1, 2, 3, 3, 5, 5},
		},
		{
			name:   "Empty list",
			input:  []int{},
			output: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quickSort(test.input)
			if !reflect.DeepEqual(test.input, test.output) {
				t.Errorf("Expected %v, but got %v", test.output, test.input)
			}
		})
	}
}
