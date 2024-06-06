package main

import "leetcode/utils"

func main() {
	console := utils.NewReaderWriter()

	lists := console.ReadKLinkedListInt()

	//lists := make([]*utils.ListNode, 2)
	//lists[0] = nil
	//lists[1] = &utils.ListNode{Val: 1}

	sorted := utils.MergeKLists(lists)
	console.PrintLinkedListInt(sorted)
}
