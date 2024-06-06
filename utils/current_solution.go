package utils

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeLists(result, lists[i])
	}
	return result
}

func mergeLists(a, b *ListNode) *ListNode {
	curr := &ListNode{}
	head := curr
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	for a != nil || b != nil {
		if a == nil {
			curr.Next = b
			return head.Next
		}
		if b == nil {
			curr.Next = a
			return head.Next
		}
		if a.Val < b.Val {
			curr.Next = a
			a = a.Next
		} else {
			curr.Next = b
			b = b.Next
		}
		curr = curr.Next
	}
	return head
}
