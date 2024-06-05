package hard

import "slices"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{Val: 0}
	prev := head
	current := &ListNode{}
	current = nil
	currentI := 0
	for len(lists) != 0 {
		i := 0
		for i < len(lists) {
			if lists[i] == nil {
				lists = slices.Delete(lists, i, i+1)
				continue
			}
			if current == nil {
				current = lists[i]
				currentI = i
			} else if lists[i].Val < current.Val {
				current = lists[i]
				currentI = i
			}
			i++
		}
		if current == nil {
			return head.Next
		}
		if current.Next == nil {
			lists = slices.Delete(lists, currentI, currentI+1)
		} else {
			lists[currentI] = lists[currentI].Next
		}
		prev.Next = current
		prev = current
		current = nil
	}
	return head.Next
}
