package main

import "leetcode/utils"

func main() {
	console := utils.NewReaderWriter()

	n := console.ReadInt()

	ans := utils.IsHappy(n)
	console.PrintBool(ans)
}
