package easy

type listNode struct {
	Val  int
	Next *listNode
}

func hasCycle(head *listNode) bool {
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
