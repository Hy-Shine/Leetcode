package go_leetcode

import (
	"testing"
)

func TestMaximumOddBinaryNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1", "1"},
		{"1101", "1101"},
		{"101010", "110001"},
	}

	for _, test := range tests {
		result := maximumOddBinaryNumber(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}
