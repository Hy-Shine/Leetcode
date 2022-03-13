package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	s := "aaaa"
	result := longestPalindrome(s)
	s2 := "aaaabbbccccc"
	result2 := longestPalindrome(s2)
	assert.Equal(t, 4, result)
	assert.Equal(t, 11, result2)
}
