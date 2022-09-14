package go_leetcode

import (
	"fmt"
	"testing"
)

func TestMyAToi(t *testing.T) {
	list := []string{
		"   -123 asd",
		"123",
		"+123",
		"  + 123",
		"+12345678901234567890",
		"-12345678901234567890 hello world",
		"+-12",
	}
	for _, v := range list {
		n := myAToi(v)
		fmt.Printf("char: %s, number: %d\n", v, n)
	}
}

func BenchmarkMyAToi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAToi("12345678901234567890")
	}
}
