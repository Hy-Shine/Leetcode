package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution1
func reverseList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	head := &ListNode{
		0,
		nil,
	}
	var temp = &ListNode{}
	var temp2 = &ListNode{}
	for l != nil {
		temp = head.Next
		head.Next = l
		temp2 = l.Next
		l.Next = temp
		l = temp2
	}
	return head.Next
}

// solution2
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
