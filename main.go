package main

import "leetcode/utils"

func main() {
	console := utils.NewReaderWriter()

	n := console.ReadInt()
	arr := console.ReadSliceInt(n)
	target := console.ReadInt()

	ans := utils.TwoSum(arr, target)
	console.PrintSlice(ans)
}
