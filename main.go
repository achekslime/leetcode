package main

import "leetcode/utils"

func main() {
	console := utils.NewReaderWriter()

	n := console.ReadString()

	ans := utils.IsValid(n)
	console.PrintBool(ans)
}
