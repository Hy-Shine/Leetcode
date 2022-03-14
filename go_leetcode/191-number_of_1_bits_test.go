package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	r := hammingWeight(4)
	assert.Equal(t, 1, r)
}

func TestHammingWeight2(t *testing.T) {
	r := hammingWeight2(4)
	assert.Equal(t, 1, r)
}
