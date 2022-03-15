package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	l := []int{9, 9, 9}
	assert.Equal(t, []int{1, 0, 0, 0}, plusOne(l))
}

func BenchmarkPlusOne(t *testing.B) {
	l := []int{9, 9, 9}
	for i := 0; i < t.N; i++ {
		plusOne(l)
	}
}

func BenchmarkPlusOne2(t *testing.B) {
	l := []int{9, 9, 9}
	for i := 0; i < t.N; i++ {
		plusOne2(l)
	}
}
