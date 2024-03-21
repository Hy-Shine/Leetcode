package go_leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-two-sorted-lists

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	var newList = &ListNode{}
	var head = newList
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			head.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 != nil {
		head.Next = &ListNode{Val: list1.Val, Next: list1.Next}
	}
	if list2 != nil {
		head.Next = &ListNode{Val: list2.Val, Next: list2.Next}
	}

	return newList.Next
}
