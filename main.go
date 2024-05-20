package main

import (
	"leetcode/utils"
)

func main() {
	console := utils.NewReaderWriter()

	// input
	n := console.ReadInt()

	// output
	ans := utils.ClimbStairs(n)
	console.PrintInt(ans)
}
