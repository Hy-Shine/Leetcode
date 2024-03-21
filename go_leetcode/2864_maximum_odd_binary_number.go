package go_leetcode

import "strings"

func maximumOddBinaryNumber(s string) string {
	cnt := strings.Count(s, "1")
	x := strings.Repeat("1", cnt-1) + strings.Repeat("0", len(s)-cnt) + "1"
	return x
}
