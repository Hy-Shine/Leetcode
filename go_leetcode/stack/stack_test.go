package stack

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	node := NewStack()

	for i := 0; i < 50; i++ {
		node.Push(i)
	}

	for !node.IsEmpty() {
		val := node.Pop()
		fmt.Println(val)
	}
}
