package stack

type nodeStack struct {
	value any
	next  *nodeStack
}

func NewStack() *nodeStack {
	return &nodeStack{}
}

func (node *nodeStack) Push(val any) {
	node.next = &nodeStack{value: val, next: node.next}
}

func (node *nodeStack) Pop() (val any) {
	if node.next != nil {
		val = node.next.value
		node.next = node.next.next
	}
	return
}

func (node *nodeStack) Value() any {
	return node.next.value
}

func (node *nodeStack) IsEmpty() bool {
	return node.next == nil
}
