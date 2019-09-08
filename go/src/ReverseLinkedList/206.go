package ReverseLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for {
		if curr == nil {
			return prev
		}

		var temp *ListNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
