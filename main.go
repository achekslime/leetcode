package main

import "leetcode/utils"

func main() {
	console := utils.NewReaderWriter()

	n := console.ReadInt()
	arr := console.ReadSliceInt(n)

	ans := utils.MissingNumber(arr)
	console.PrintInt(ans)
}
