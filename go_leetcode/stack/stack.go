package stack

type nodeStack struct {
	value int
	next  *nodeStack
}

func NewStack() *nodeStack {
	return &nodeStack{}
}

func (node *nodeStack) Push(val int) {
	node.next = &nodeStack{value: val, next: node.next}
}

func (node *nodeStack) Pop() (val int) {
	if node.next != nil {
		val = node.next.value
		node.next = node.next.next
	}
	return
}

func (node *nodeStack) Value() int {
	return node.next.value
}

func (node *nodeStack) IsEmpty() bool {
	return node.next == nil
}

func (node *nodeStack) Len() (length int) {
	for head := node; head.next != nil; head = head.next {
		length++
	}
	return length
}
