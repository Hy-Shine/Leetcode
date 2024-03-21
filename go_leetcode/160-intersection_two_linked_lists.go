package go_leetcode

// level: Easy
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[*ListNode]struct{})
	for headA != nil || headB != nil {
		if _, ok := m[headA]; ok && headA != nil {
			return headA
		}
		if headA != nil {
			m[headA] = struct{}{}
			headA = headA.Next
		}

		if _, ok := m[headB]; ok && headB != nil {
			return headB
		}
		if headB != nil {
			m[headB] = struct{}{}
			headB = headB.Next
		}
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	la, lb := headA, headB
	for la != lb {
		if la != nil {
			la = la.Next
		} else {
			la = headB
		}

		if lb != nil {
			lb = lb.Next
		} else {
			lb = headA
		}
	}

	return la
}
