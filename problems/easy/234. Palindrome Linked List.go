package easy

func isPalindrome(head *ListNode) bool {
	slice := make([]int, 0)
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	l := len(slice)
	for i := 0; i < l/2; i++ {
		if slice[i] != slice[l-i-1] {
			return false
		}
	}
	return true
}
