package go_leetcode

import (
	"strconv"
)

func reverse(x int) int {
	s := strconv.Itoa(x)
	s1 := 1
	if s[0] == '-' {
		s1 = -1
		s = s[1:]
	}

	strByte := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		strByte[i], strByte[len(s)-1-i] = strByte[len(s)-1-i], strByte[i]
	}
	strInt, _ := strconv.Atoi(string(strByte))

	if s1 < 0 && strInt > 1<<31 {
		strInt = 0
	}
	if s1 > 0 && strInt > 1<<31-1 {
		strInt = 0
	}

	return strInt * s1
}
