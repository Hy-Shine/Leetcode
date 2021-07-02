package main

import "fmt"

func main() {
	s := "hello"
	s0 := "ello"
	fmt.Println(strStr(s, s0))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	len1, len2 := len(haystack), len(needle)
	for i := 0; i <= len1-len2; i++ {
		if haystack[i:len2+i] == needle {
			return i
		}
	}
	return -1
}
