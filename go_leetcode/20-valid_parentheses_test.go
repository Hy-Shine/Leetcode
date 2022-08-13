package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	r1 := "()[]{}"
	r2 := "{[]}()"
	r3 := "(){}[(])"
	assert.Equal(t, true, IsValidParentheses(r1))
	assert.Equal(t, true, IsValidParentheses(r2))
	assert.Equal(t, false, IsValidParentheses(r3))
}
