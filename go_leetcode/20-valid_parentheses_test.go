package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	r1 := "()[]{}"
	r2 := "{[]}()"
	assert.Equal(t, true, isValid(r1))
	assert.Equal(t, true, isValid(r2))
}
