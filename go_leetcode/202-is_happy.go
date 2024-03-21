package go_leetcode

import (
	"strconv"
	"strings"
)

func isHappy(n int) bool {
	var val string
	var onlyOneAndZero bool
	m := make(map[int]struct{})
	for {
		var x int
		for n > 0 {
			x += (n % 10) * (n % 10)
			n = n / 10
		}

		val = strconv.Itoa(x)
		onlyOneAndZero = strings.Count(val, "1") == 1 && strings.Count(val, "0") == len(val)-1
		if onlyOneAndZero {
			return true
		}

		if _, ok := m[x]; ok {
			return false
		}
		m[x] = struct{}{}
		x = n
	}
}

func mergeAlternately(word1 string, word2 string) string {
	word1Len := len(word1)
	word2Len := len(word2)

	var builder strings.Builder
	builder.Grow(word1Len + word2Len)
	for i := 0; i < word1Len || i < word2Len; i++ {
		if i < word1Len {
			builder.WriteByte(word1[i])
		}
		if i < word2Len {
			builder.WriteByte(word2[i])
		}
	}

	return builder.String()
}
