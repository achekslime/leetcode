package main

import "leetcode/utils"

func main() {
	console := utils.NewReaderWriter()

	count := console.ReadInt()
	lists := console.ReadKLinkedListInt(count)

	//lists := make([]*utils.ListNode, 2)
	//lists[0] = nil
	//lists[1] = &utils.ListNode{Val: 1}

	sorted := utils.MergeKLists(lists)
	console.PrintLinkedListInt(sorted)
}
