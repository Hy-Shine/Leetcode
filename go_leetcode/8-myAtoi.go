package go_leetcode

import (
	"bytes"
	"strconv"
)

func myAToi(s string) int {
	builder := bytes.Buffer{}
	flag := 1
	allocFlag := false
	for i := 0; i < len(s); i++ {
		if (s[i] == ' ' || s[i] == '-' || s[i] == '+') && !allocFlag {
			if s[i] == '-' || s[i] == '+' {
				allocFlag = true
				if s[i] == '-' {
					flag = -1
				}
			}
			continue
		}
		if !isValidNumber(s[i]) {
			break
		}
		allocFlag = true
		builder.WriteByte(s[i])
	}

	str2Int, _ := strconv.Atoi(builder.String())
	return checkValue(str2Int * flag)
}

func isValidNumber(a byte) bool {
	return '0' <= a && a <= '9'
}

func checkValue(n int) int {
	const limit = 1 << 31
	if n < -1*limit {
		n = -1 * limit

	}
	if n > limit-1 {
		n = limit - 1
	}
	return n
}
