package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1, len2 := getLinkLength(l1), getLinkLength(l2)
	if len1 >= len2 {
		return handler(l1, l2)
	}
	return handler(l2, l1)
}

func getLinkLength(l *ListNode) int {
	length := 0
	for l != nil {
		length++
		l = l.Next
	}
	return length
}

func handler(l1 *ListNode, l2 *ListNode) *ListNode {
	head, sum := l1, 0
	for l1 != nil {
		if l2 != nil {
			sum = l1.Val + l2.Val
			if sum >= 10 {
				l1.Val = sum - 10
				if l1.Next != nil {
					l1.Next.Val = l1.Next.Val + 1
				} else {
					l1.Next = &ListNode{
						1,
						nil,
					}
				}
			} else {
				l1.Val = sum
			}
			//
			if l2.Next != nil {
				l2 = l2.Next
			} else {
				l2 = nil
			}
		} else {
			if l1.Val >= 10 {
				l1.Val = l1.Val - 10
				if l1.Next == nil {
					l1.Next = &ListNode{
						1,
						nil,
					}
				} else {
					l1.Next.Val = l1.Next.Val + 1
				}
			}
		}
		l1 = l1.Next
	}
	return head
}
