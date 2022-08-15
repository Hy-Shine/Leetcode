package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func NewLinkedList() *Node {
	return &Node{}
}

func (node *Node) Append(values ...int) {
	head := node
	for head.Next != nil {
		head = head.Next
	}
	valueNode := NewLinkedList()
	for _, v := range values {
		valueNode.Next = &Node{Value: v, Next: valueNode.Next}
	}
	valueNode.Reverse()
	head.Next = valueNode.Next
}

func (node *Node) AppendHead(values ...int) {
	for _, v := range values {
		node.Next = &Node{Value: v, Next: node.Next}
	}
}

func (node *Node) Len() (length int) {
	head := node
	for head.Next != nil {
		length++
		head = head.Next
	}
	return
}

func (node *Node) Delete(val int) {
	head := node
	for head.Next != nil {
		if head.Next.Value == val {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
}

func (node *Node) Clear() {
	node.Next = nil
}

func (node *Node) MergeNode(n *Node) {
	head := node
	for head.Next != nil {
		head = head.Next
	}
	head.Next = n.Next
}

func (node *Node) Range() {
	head := node
	for head.Next != nil {
		fmt.Println(head.Next.Value)
		head = head.Next
	}
}

func (node *Node) Reverse() {
	node2 := NewLinkedList()
	head := node
	for head.Next != nil {
		node2.Next = &Node{
			Value: head.Next.Value,
			Next:  node2.Next,
		}
		head = head.Next
	}
	node.Next = node2.Next
}
