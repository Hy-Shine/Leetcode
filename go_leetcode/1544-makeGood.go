package go_leetcode

type nodeStack struct {
	value rune
	next  *nodeStack
}

func (node *nodeStack) Push(val rune) {
	node.next = &nodeStack{value: val, next: node.next}
}

func (node *nodeStack) Pop() (val rune) {
	if node.next != nil {
		val = node.next.value
		node.next = node.next.next
	}
	return
}

func (node *nodeStack) Value() rune {
	if !node.IsEmpty() {
		return node.next.value
	}
	return 0
}

func (node *nodeStack) IsEmpty() bool {
	return node.next == nil
}

func MakeGood(s string) string {
	node := &nodeStack{}
	const d = 'a' - 'A'
	for _, v := range s {
		if !node.IsEmpty() && (node.Value()-v == d || v-node.Value() == d) {
			node.Pop()
			continue
		}
		node.Push(v)
	}

	str := make([]rune, 0)
	for !node.IsEmpty() {
		str = append(str, node.Pop())
	}
	str2 := make([]rune, 0, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		str2 = append(str2, str[i])
	}
	return string(str2)
}
