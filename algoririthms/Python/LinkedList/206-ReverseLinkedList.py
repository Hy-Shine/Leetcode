# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head == None:
            return head
        v_head = ListNode(None)
        v_head.next = head
        while head.next != None:
            cur = head.next
            head.next = head.next.next
            cur.next = v_head.next
            v_head.next = cur
        return v_head.next


# testing cases
head = ListNode(1)
b = ListNode(2)
c = ListNode(3)
d = ListNode(4)
head.next = b
b.next = c
c.next = d
L = Solution()
s = L.reverseList(head)
while s != None:
    print(s.val)
    s = s.next
