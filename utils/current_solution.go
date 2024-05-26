package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head.Next
	parent := head
	head.Next = nil
	for current.Next != nil {
		oldNext := current.Next
		current.Next = parent
		parent = current
		current = oldNext
	}
	current.Next = parent
	return current
}
