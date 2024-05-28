package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	if slow == nil || fast == nil {
		return false
	}
	for fast.Next != nil && slow.Next != nil {
		fast = fast.Next
		if fast == slow {
			return true
		}
		if fast.Next == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
		slow = slow.Next
	}

	return false
}
