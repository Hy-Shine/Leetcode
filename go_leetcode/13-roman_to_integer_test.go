package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	s := "MCMXCIV"
	r := romanToInt(s)
	assert.Equal(t, 1994, r)
}
