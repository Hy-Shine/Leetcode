package linkedlist

import "testing"

func TestReverseList(t *testing.T) {
	node1 := NewLinkedList()
	for i := 0; i < 50; i++ {
		node1.Append(i)
	}

	node1.Reverse()
	node1.Range()
}
