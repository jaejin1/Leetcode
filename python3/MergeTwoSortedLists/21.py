class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        dummy = cur = ListNode(0)
        print(id(dummy))
        print(id(cur))

        while l1 and l2:
            if l1.val < l2.val:
                cur.next = ListNode(l1.val)
                # cur.next = l1
                l1 = l1.next
            else:
                cur.next = ListNode(l2.val)
                # cur.next = l2
                l2 = l2.next
            
            cur = cur.next
            
        # 남은거 정리
        cur.next = l1 or l2
        return dummy.next
