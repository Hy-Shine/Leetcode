package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	list1 := []string{"abc"}
	result1 := longestCommonPrefix(list1)

	list2 := []string{"abc", "ab", "alldj"}
	result2 := longestCommonPrefix(list2)

	assert.Equal(t, "abc", result1)
	assert.Equal(t, "a", result2)
}
