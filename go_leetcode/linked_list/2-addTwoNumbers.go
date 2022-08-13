package linkedlist

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	len1, len2 := getLinkLength(l1), getLinkLength(l2)
	if len1 >= len2 {
		return handler(l1, l2)
	}
	return handler(l2, l1)
}

func getLinkLength(l *Node) int {
	length := 0
	for l != nil {
		length++
		l = l.Next
	}
	return length
}

func handler(l1 *Node, l2 *Node) *Node {
	head, sum := l1, 0
	for l1 != nil {
		if l2 != nil {
			sum = l1.Value + l2.Value
			if sum >= 10 {
				l1.Value = sum - 10
				if l1.Next != nil {
					l1.Next.Value = l1.Next.Value + 1
				} else {
					l1.Next = &Node{
						1,
						nil,
					}
				}
			} else {
				l1.Value = sum
			}
			//
			if l2.Next != nil {
				l2 = l2.Next
			} else {
				l2 = nil
			}
		} else {
			if l1.Value >= 10 {
				l1.Value = l1.Value - 10
				if l1.Next == nil {
					l1.Next = &Node{
						1,
						nil,
					}
				} else {
					l1.Next.Value = l1.Next.Value + 1
				}
			}
		}
		l1 = l1.Next
	}
	return head
}
