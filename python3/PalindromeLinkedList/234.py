class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        fast = slow = head
        
        # slow is half of head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        
        node = None
        # node is reverse Linked List of slow
        
        while slow:
            nxt = slow.next
            slow.next = node
            node = slow
            slow = nxt
        
        while node:
            if head.val != node.val:
                return False
            head = head.next
            node = node.next
    
        return True