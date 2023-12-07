package go_leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	// Testing when nums is empty
	nums := []int{}
	val := 1
	expected := 0
	result := removeElement(nums, val)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	// Testing when nums has only one element and it matches val
	nums = []int{1}
	val = 1
	expected = 0
	result = removeElement(nums, val)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	// Testing when nums has only one element and it does not match val
	nums = []int{2}
	val = 1
	expected = 1
	result = removeElement(nums, val)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	// Testing when nums has multiple elements and some match val
	nums = []int{0, 1, 2, 3, 4, 5, 1, 3}
	val = 1
	expected = 6
	result = removeElement(nums, val)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	// Testing when nums has multiple elements and none match val
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7}
	val = 8
	expected = 8
	result = removeElement(nums, val)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
