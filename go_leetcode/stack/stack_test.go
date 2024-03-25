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

func TestStacks(t *testing.T) {
	stack := NewStackS[int]()
	var count int
	for ; count < 100; count++ {
		stack.Push(count)
	}

	for !stack.IsEmpty() {
		count--
		pop := stack.Pop()
		if pop != count {
			t.Errorf("want %d, but got %d\n", count, pop)
		}
	}
}
