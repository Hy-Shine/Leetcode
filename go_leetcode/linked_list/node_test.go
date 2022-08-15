package linkedlist

import (
	"fmt"
	"testing"
)

func TestNodeLen(t *testing.T) {
	node := NewLinkedList()
	node.Next = &Node{Value: 10}
	l := node.Len()
	fmt.Println(l)
}

func TestNodeAppend(t *testing.T) {
	node := NewLinkedList()
	for i := 0; i < 100; i++ {
		node.Append(i)
	}
	fmt.Println("length:", node.Len())

	node.Delete(0)
	node.Range()
}

func TestNodeAppendNode(t *testing.T) {
	node1 := NewLinkedList()
	node2 := NewLinkedList()
	for i := 0; i < 100; i += 2 {
		node1.Append(i)
	}
	for i := 1; i < 100; i += 2 {
		node2.Append(i)
	}

	node1.MergeNode(node2)
	node1.Reverse()
	node1.Range()
}

func BenchmarkNodeAppend(b *testing.B) {
	node := NewLinkedList()
	for i := 0; i < 1000000; i++ {
		node.AppendHead(i)
	}
}
