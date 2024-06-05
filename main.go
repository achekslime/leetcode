package main

import "leetcode/utils"

func main() {
	console := utils.NewReaderWriter()

	root := &utils.TreeNode{Val: 3}

	isSym := utils.IsSymmetric(root)
	console.PrintBool(isSym)
}
