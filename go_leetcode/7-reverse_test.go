package go_leetcode

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	list := []int{
		12000,
		123,
		-1234,
		-0,
		0,
		-12000,
	}
	for _, v := range list {
		fmt.Println(v, reverse(v))
	}
}
